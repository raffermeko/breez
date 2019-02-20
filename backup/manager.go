package backup

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/breez/breez/config"
	"github.com/breez/breez/data"
	"github.com/breez/breez/db"
	"github.com/breez/lightninglib/lnrpc"
)

var (
	backupDelay = time.Duration(time.Second * 2)
)

/*
RequestBackup push a request for the backup files of breez
*/
func (b *Manager) RequestBackup() {

	// first thing push a pending backup request to the database so we
	// can recover in case of error.
	err := b.db.addBackupRequest()
	if err != nil {
		log.Errorf("failed to set pending backup %v", err)
		b.notifyBackupFailed(err)
		return
	}

	go func() {
		select {
		case <-time.After(backupDelay):
		case <-b.quitChan:
			return
		}
		b.backupRequestChan <- struct{}{}
	}()
}

// Restore handles all the restoring process:
// 1. Downloading the backed up files for a specific node id.
// 2. Put the backed up files in the right place according to the configuration
func (b *Manager) Restore(nodeID string) error {
	backupID, err := b.getBackupIdentifier()
	if err != nil {
		return err
	}
	files, err := b.provider.DownloadBackupFiles(nodeID, backupID)
	if err != nil {
		return err
	}
	if len(files) != 3 {
		return fmt.Errorf("wrong number of backup files %v", len(files))
	}
	c, err := config.GetConfig(b.workingDir)
	if err != nil {
		return err
	}

	paths := map[string]string{
		"wallet.db":  "data/chain/bitcoin/{{network}}",
		"channel.db": "data/graph/{{network}}",
		"breez.db":   "",
	}
	for _, f := range files {
		basename := path.Base(f)
		p, ok := paths[basename]
		if !ok {
			return err
		}
		destDir := path.Join(b.workingDir, strings.Replace(p, "{{network}}", c.Network, -1))
		if destDir != b.workingDir {
			err = os.MkdirAll(destDir, 0700)
			if err != nil {
				return err
			}
		}
		err = os.Rename(f, path.Join(destDir, basename))
		if err != nil {
			return err
		}
	}
	return nil
}

// AvailableSnapshots returns a list of snsapshot that the backup provider reports
// it has. Every snapshot is for a specific node id.
func (b *Manager) AvailableSnapshots() ([]SnapshotInfo, error) {
	return b.provider.AvailableSnapshots()
}

// IsSafeToRunNode checks if it is safe for this breez instance to run a specific node.
// It is considered safe if we don't know of another instance which is the last to restore
// this node (nodeID)
func (b *Manager) IsSafeToRunNode(nodeID string) (bool, error) {
	snapshots, err := b.provider.AvailableSnapshots()
	if err != nil {
		return false, err
	}
	backupID, err := b.getBackupIdentifier()
	if err != nil {
		return false, err
	}
	for _, s := range snapshots {
		if s.NodeID == nodeID && s.BackupID != "" && backupID != s.BackupID {
			log.Errorf("remote restore was found for node %v.", nodeID)
			log.Errorf("current backupID=%v, remote backupID-%v", backupID, s.BackupID)
			return false, nil
		}
	}
	return true, nil
}

// Start is the main go routine that listens to backup requests and is resopnsible for executing it.
func (b *Manager) Start(lightningClient lnrpc.LightningClient, breezDB *db.DB) {
	if atomic.SwapInt32(&b.started, 1) == 1 {
		return
	}
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()

		for {
			select {
			case <-b.backupRequestChan:
				//First get the last pending request in the database
				pendingID, err := b.db.lastBackupRequest()
				if pendingID == 0 {
					continue
				}

				paths, nodeID, err := b.prepareBackupInfo(lightningClient, breezDB)
				if err != nil {
					log.Errorf("error in backup %v", err)
					b.notifyBackupFailed(err)
					continue
				}
				if err := b.provider.UploadBackupFiles(paths, nodeID); err != nil {
					log.Errorf("error in backup %v", err)
					b.notifyBackupFailed(err)
					continue
				}
				b.db.markBackupRequestCompleted(pendingID)
				log.Infof("backup finished succesfully")
				b.ntfnChan <- data.NotificationEvent{Type: data.NotificationEvent_BACKUP_SUCCESS}
			case <-b.quitChan:
				return
			}
		}
	}()

	// execute recovery if needed.
	b.backupRequestChan <- struct{}{}
}

// Stop stops the BackupService and wait for complete shutdown.
func (b *Manager) Stop() {
	if atomic.SwapInt32(&b.stopped, 1) == 1 {
		return
	}

	close(b.quitChan)
	b.wg.Wait()
}

func (b *Manager) notifyBackupFailed(err error) {
	log.Infof("notifyBackupFailed %v", err)
	if ferr, ok := err.(ProviderError); ok {
		if ferr.IsAuthError() {
			b.ntfnChan <- data.NotificationEvent{Type: data.NotificationEvent_BACKUP_AUTH_FAILED}
			return
		}
	}
	b.ntfnChan <- data.NotificationEvent{Type: data.NotificationEvent_BACKUP_FAILED}
}

func (b *Manager) breezdbCopy(breezDB *db.DB) (string, error) {
	dir, err := ioutil.TempDir("", "backup")
	if err != nil {
		return "", err
	}
	return breezDB.BackupDb(dir)
}

// extractBackupInfo extracts the information that is needed for the external backup service:
// 1. paths - the files need to be backed up.
// 2. nodeID - the current lightning node id.
// 3. backupID - an identifier for this instance of breez. It is needed for conflict detection.
func (b *Manager) prepareBackupInfo(lightningClient lnrpc.LightningClient, breezDB *db.DB) (paths []string, nodeID string, err error) {
	log.Infof("extractBackupInfo started")
	response, err := lightningClient.GetBackup(context.Background(), &lnrpc.GetBackupRequest{})
	if err != nil {
		log.Errorf("Couldn't get backup: %v", err)
		return nil, "", err
	}
	info, err := lightningClient.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
	if err != nil {
		return nil, "", err
	}

	f, err := b.breezdbCopy(breezDB)
	if err != nil {
		log.Errorf("Couldn't get breez backup file: %v", err)
		return nil, "", err
	}
	files := append(response.Files, f)
	log.Infof("extractBackupInfo completd")
	return files, info.IdentityPubkey, nil
}

// getBackupIdentifier retrieves an identifier that is unique for this instance of breez.
// We use is as a mechanism for conflict detection, in case a restore was done on some
// other device.
// The identifier is generated once and save in a file, we can't save it in the db as it
// is backed up and would cause the restored node to have the same identifier...
func (b *Manager) getBackupIdentifier() (string, error) {
	backupDir := path.Join(b.workingDir, "backup")
	if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
		return "", err
	}
	backupFile := path.Join(backupDir, "breez_backup_id")
	if _, err := os.Stat(backupFile); os.IsNotExist(err) {
		var id [32]byte
		_, err = rand.Read(id[:])
		if err != nil {
			return "", err
		}
		if err = ioutil.WriteFile(backupFile, id[:], os.ModePerm); err != nil {
			return "", err
		}
	}
	id, err := ioutil.ReadFile(backupFile)
	if err != nil {
		return "", err
	}
	return "backup-id-" + hex.EncodeToString(id), nil
}
