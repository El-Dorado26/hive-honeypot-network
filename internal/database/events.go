package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

// DBStore provides database operations for HIVE.
type DBStore struct {
	conn *sql.DB
}

// NewDBStore creates a store using an existing database connection.
func NewDBStore(conn *sql.DB) *DBStore {
	return &DBStore{conn: conn}
}

// CreateEvent inserts one event into PostgreSQL.
func (db *DBStore) CreateEvent(
	ctx context.Context,
	source string,
	eventType string,
	occurredAt time.Time,
	eventData []byte,
) (int64, time.Time, error) {
	var id int64
	var receivedAt time.Time

	err := db.conn.QueryRowContext(
		ctx,
		`INSERT INTO events (source, event_type, occurred_at, event_data)
		 VALUES ($1, $2, $3, $4::jsonb)
		 RETURNING id, received_at`,
		source,
		eventType,
		occurredAt,
		json.RawMessage(eventData),
	).Scan(&id, &receivedAt)

	if err != nil {
		return 0, time.Time{}, fmt.Errorf("insert event: %w", err)
	}

	return id, receivedAt, nil
}

// Event represents one stored HIVE event.
type Event struct {
	ID         int64           `json:"id"`
	Source     string          `json:"source"`
	EventType  string          `json:"event_type"`
	OccurredAt time.Time       `json:"occurred_at"`
	EventData  json.RawMessage `json:"event_data"`
	ReceivedAt time.Time       `json:"received_at"`
}

// ListEvents retrieves events in newest first order
func (db *DBStore) ListEvents(
	ctx context.Context,
	limit int,
	offset int,
) ([]Event, error) {
	rows, err := db.conn.QueryContext(
		ctx,
		`SELECT id, source, event_type, occurred_at, event_data, received_at
		 FROM events
		 ORDER BY occurred_at DESC, id DESC
		 LIMIT $1 OFFSET $2`,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}
	defer rows.Close()

	events := make([]Event, 0, limit)

	for rows.Next() {
		var event Event
		var eventData []byte

		err := rows.Scan(
			&event.ID,
			&event.Source,
			&event.EventType,
			&event.OccurredAt,
			&eventData,
			&event.ReceivedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}

		event.EventData = json.RawMessage(eventData)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate events: %w", err)
	}

	return events, nil
}
