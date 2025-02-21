package service

import "log/slog"

func (s *service) NotificationWorker() {
	for notify := range s.SendNotification {
		slog.Info("Send Notification",
			"Card:", notify.Card, "OrderType:", notify.OrderType, "SessionID:", notify.SessionID, "WebsiteUrl:", notify.WebsiteUrl)
	}
}

func (s *service) CloseNotificationWorker() {
	slog.Info("Close notification worker channel")
	close(s.SendNotification)
}
