package service

import (
	"AyanTestProject/internal/model"
	"testing"
	"time"
)

func TestNotificationWorker(t *testing.T) {
	s := &service{
		SendNotification: make(chan *model.Event, 1),
	}

	go s.NotificationWorker()

	testNotification := &model.Event{
		Card:       "4433**1409",
		OrderType:  "Purchase",
		SessionID:  "test-session",
		WebsiteUrl: "https://example.com",
	}

	s.SendNotification <- testNotification

	time.Sleep(100 * time.Millisecond)

	s.CloseNotificationWorker()
}
