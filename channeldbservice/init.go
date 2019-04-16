package channeldbservice

import (
	"path"
	"strings"

	"github.com/breez/breez/config"
	"github.com/breez/breez/log"
	"github.com/breez/breez/refcount"
	"github.com/breez/lightninglib/channeldb"
	"github.com/btcsuite/btclog"
)

const (
	directoryPattern = "data/graph/{{network}}/"
)

var (
	serviceRefCounter refcount.ReferenceCountable
	chanDB            *channeldb.DB
	logger            btclog.Logger
)

// Get returns a Ch
func Get(workingDir string) (db *channeldb.DB, cleanupFn func() error, err error) {
	service, release, err := serviceRefCounter.Get(
		func() (interface{}, refcount.ReleaseFunc, error) {
			return newService(workingDir)
		},
	)
	return service.(*channeldb.DB), release, err
}

func newService(workingDir string) (*channeldb.DB, refcount.ReleaseFunc, error) {
	db, err := createService(workingDir)
	if err != nil {
		return nil, nil, err
	}
	return db, release, err
}

func release() error {
	return chanDB.Close()
}

func createService(workingDir string) (*channeldb.DB, error) {
	config, err := config.GetConfig(workingDir)
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logBackend, err := log.GetLogBackend(workingDir)
		if err != nil {
			return nil, err
		}
		logger = logBackend.Logger("CHANNELDB")
		logger.SetLevel(btclog.LevelDebug)
	}
	logger.Infof("creating shared channeldb service.")
	graphDir := path.Join(workingDir, strings.Replace(directoryPattern, "{{network}}", config.Network, -1))
	chanDB, err := channeldb.Open(graphDir)
	if err != nil {
		logger.Errorf("unable to open channeldb: %v", err)
		return nil, err
	}

	logger.Infof("channeldb was opened successfuly")
	return chanDB, err
}
