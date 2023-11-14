package components

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB = ConnectPlanetscale()

func ConnectPlanetscale() *sql.DB {
	dsn := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	log.Println("Successfully connected to PlanetScale!")
	return db
}
