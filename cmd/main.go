package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"static-site-hosting/handlers"
	"static-site-hosting/middleware"
	"static-site-hosting/routes"
)

func main() {

	if err := os.MkdirAll("sites", os.ModePerm); err != nil {
		log.Fatalf("Failed to create sites folder: %v", err)
	}
	// Setup and connect to the database
	if err := os.Remove("./db/database.db"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error removing database file: %v", err)
	}

	db, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create a sample table for demonstration purposes
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS example (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", routes.HelloWorldHandler)
	mux.HandleFunc("/deploy", handlers.UploadHandler)
	mux.Handle("/sites/", http.StripPrefix("/sites/", http.FileServer(http.Dir("./sites"))))
	mux.HandleFunc("/deployments", handlers.ListDeployments)

	wrappedMux := middleware.LoggingMiddleware(mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))

}
