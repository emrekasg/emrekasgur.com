package components

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB = ConnectPlanetscale()

func ConnectPlanetscale() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	log.Println("Successfully connected to database!")
	return db
}
