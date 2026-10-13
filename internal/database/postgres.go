
package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open connects to PostgreSQL and verifies that the database is reachable.
func Open() (*sql.DB, error) {
	password := os.Getenv("HIVE_DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("HIVE_DB_PASSWORD environment variable is not set")
	}

	host := os.Getenv("HIVE_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("HIVE_DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("HIVE_DB_USER")
	if user == "" {
		user = "hive_app"
	}

	dbName := os.Getenv("HIVE_DB_NAME")
	if dbName == "" {
		dbName = "hive_db"
	}

	sslMode := os.Getenv("HIVE_DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   dbName,
	}

	query := url.Values{}
	query.Set("sslmode", sslMode)
	dsn.RawQuery = query.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("configure PostgreSQL connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("connect to PostgreSQL: %w", err)
	}

	return db, nil
}
