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
