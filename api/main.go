package main

import (
	"fmt"
	"log"

	"github.com/emrekasg/personal-website-api/components"
	"github.com/emrekasg/personal-website-api/webserver"
)

func main() {
	fmt.Println("Starting the application...")

	// Initialize database connection
	db, err := components.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	fmt.Println("Database connection established")
	fmt.Println("Starting the server...")
	webserver.RunApp()
}
