package components

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	initialDelay := time.Second
	maxDelay := 30 * time.Minute
	retryDelay := initialDelay
	attempt := 1

	for {
		db, err := sql.Open("postgres", os.Getenv("DSN"))
		if err != nil {
			log.Error().
				Err(err).
				Int("attempt", attempt).
				Msg("Failed to open database connection")

			time.Sleep(retryDelay)
			retryDelay *= 2
			if retryDelay > maxDelay {
				retryDelay = maxDelay
			}
			attempt++
			continue
		}

		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(15 * time.Minute)

		// Verify connection
		if err := db.Ping(); err != nil {
			db.Close()
			log.Error().
				Err(err).
				Int("attempt", attempt).
				Msg("Failed to ping database")

			time.Sleep(retryDelay)
			retryDelay *= 2 // Exponential backoff
			if retryDelay > maxDelay {
				retryDelay = maxDelay
			}
			attempt++
			continue
		}

		DB = db
		log.Info().
			Int("attempt", attempt).
			Int("maxOpenConns", 25).
			Int("maxIdleConns", 25).
			Dur("connMaxLifetime", 15*time.Minute).
			Msg("Successfully connected to database")
		return db, nil
	}
}
