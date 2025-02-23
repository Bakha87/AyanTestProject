package repository

import (
	"AyanTestProject/internal/model"
	"context"
	"fmt"
	"time"
)

func (r *repository) PutEventNotification(ctx context.Context, req *model.Event) error {
	query := `INSERT INTO ayan.public.events(order_type, session_id, card, event_date, website_url) 
 			VALUES ($1, $2, $3, $4, $5)`

	_, err := r.pg.Pool.Exec(ctx, query, req.OrderType, req.SessionID, req.Card, req.EventDate, req.WebsiteUrl)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
