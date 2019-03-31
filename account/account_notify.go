package account

import (
	"strings"
	"time"

	breezservice "github.com/breez/breez/breez"
)

var (
	notificationTypeChannelOpened        = 1
	notificationTypeReceivePaymentsReady = 2
)

//RegisterReceivePaymentReadyNotification register in breez server for notification regarding the confirmation
//of an existing pending channel. If there is not channels and no pending channels then this function waits for
//for a channel to be opened.
func (a *Service) RegisterReceivePaymentReadyNotification(token string) error {
	return a.setUserNotificationRequest(token, notificationTypeReceivePaymentsReady)
}

//RegisterChannelOpenedNotification register in breez server for notification regarding the confirmation
//of an existing pending channel. If there is not channels and no pending channels then this function waits for
//for a channel to be opened.
func (a *Service) RegisterChannelOpenedNotification(token string) error {
	a.log.Infof("RegisterChannelOpenedNotification called")
	return a.setUserNotificationRequest(token, notificationTypeChannelOpened)
}

func (a *Service) setUserNotificationRequest(token string, notificationType int) error {
	a.log.Infof("setUserNotificationRequest notificationType = %v", notificationType)
	channelPoints, _, err := a.getBreezOpenChannels()
	if err != nil {
		return err
	}
	if len(channelPoints) > 0 {
		return nil
	}

	a.subscriptionsSync.Lock()
	a.notification = &notificationRequest{
		token:            token,
		notificationType: notificationType}
	a.subscriptionsSync.Unlock()
	return a.registerPendingChannelConfirmation()
}

func (a *Service) registerPendingChannelConfirmation() error {
	a.log.Infof("registerPendingChannelConfirmation checking for pending channel")
	a.subscriptionsSync.Lock()
	currentRequest := a.notification
	a.subscriptionsSync.Unlock()
	if currentRequest == nil {
		a.log.Infof("registerPendingChannelConfirmation not request to process")
		return nil
	}

	pendingChannelPoint, err := a.getPendingBreezChannelPoint()
	if err != nil {
		a.log.Infof("registerPendingChannelConfirmation error in querying for pending channels %v", err)
		return err
	}
	if pendingChannelPoint == "" {
		a.log.Infof("registerPendingChannelConfirmation no pending channel found")
		return nil
	}

	c, ctx, cancel := a.breezServices.NewFundManager()
	defer cancel()

	a.log.Infof("registerPendingChannelConfirmation for token %v and notification type = %v", currentRequest.token, currentRequest.notificationType)

	notificationTypeNeeded := breezservice.RegisterTransactionConfirmationRequest_READY_RECEIVE_PAYMENT
	if currentRequest.notificationType == notificationTypeChannelOpened {
		notificationTypeNeeded = breezservice.RegisterTransactionConfirmationRequest_CHANNEL_OPENED
	}
	_, err = c.RegisterTransactionConfirmation(ctx,
		&breezservice.RegisterTransactionConfirmationRequest{
			NotificationToken: currentRequest.token,
			TxID:              strings.Split(pendingChannelPoint, ":")[0],
			NotificationType:  notificationTypeNeeded,
		})
	if err != nil {
		a.subscriptionsSync.Lock()
		a.notification = nil
		a.subscriptionsSync.Unlock()
	}
	return err
}

func (a *Service) onRoutingNodePendingChannel() {
	a.registerPendingChannelConfirmation()
	a.onAccountChanged()
}

func (a *Service) onRoutingNodeOpenedChannel() {
	a.onAccountChanged()
}

func (a *Service) trackOpenedChannel() {
	defer a.wg.Done()
	defer func() {
		a.log.Info("trackOpenedChannel stopped")
	}()
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			if !a.daemonRPCReady() {
				return
			}
			channelPoints, _, err := a.getBreezOpenChannels()
			if err == nil && len(channelPoints) > 0 {
				ticker.Stop()
				a.onRoutingNodeOpenedChannel()
				return
			}
		case <-a.quitChan:
			ticker.Stop()
			return
		}
	}
}
