package service

import (
	"AyanTestProject/internal/model"
	"context"
	"fmt"
)

func (s *service) EventNotification(ctx context.Context, event *model.Event) error {
	if err := s.repository.PutEventNotification(ctx, event); err != nil {
		return fmt.Errorf("EventNotification: %w", err)
	}

	s.SendNotification <- event

	return nil
}
