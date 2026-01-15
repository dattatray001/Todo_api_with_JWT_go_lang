package database

import (
	"context" // Used to control request lifetime (timeouts, cancellation)
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect initializes and returns a PostgreSQL connection pool
// databaseURL example:
// postgres://user:password@localhost:5432/dbname?sslmode=disable
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	// Connect initializes and returns a PostgreSQL connection pool
	// databaseURL example:
	// postgres://user:password@localhost:5432/dbname?sslmode=disable
	var ctx context.Context = context.Background()

	// pgxpool.Config holds configuration parsed from DATABASE_URL
	var config *pgxpool.Config
	var err error

	// Parse DATABASE_URL into pgxpool configuration
	// This validates the connection string format
	config, err = pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Unable to parse DATABASE_URL: %v", err)
		return nil, err
	}

	// Create a new PostgreSQL connection pool using the parsed config
	// The pool manages multiple DB connections efficiently
	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool: %v", err)
		return nil, err
	}

	// Ping the database to verify that the connection is actually working
	// This helps fail fast if DB is down or credentials are wrong
	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database: %v", err)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database")
	return pool, nil
}
