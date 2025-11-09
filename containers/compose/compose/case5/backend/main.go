package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Paste struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var db *pgxpool.Pool

func main() {
	var err error
	dbURL := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	// Retry connecting to the DB with backoff
	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = pgxpool.New(context.Background(), dbURL)
		if err == nil {
			// attempt a short ping to verify connectivity
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			err = db.Ping(ctx)
			cancel()
		}
		if err == nil {
			log.Printf("Connected to DB on attempt %d\n", attempt)
			break
		}
		log.Printf("Attempt %d: failed to connect to DB: %v", attempt, err)
		// simple linear backoff (increase wait each attempt)
		time.Sleep(time.Duration(attempt) * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to DB after %d attempts: %v", maxAttempts, err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS pastes (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/pastes", handlePastes)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePastes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "GET" {
		rows, err := db.Query(context.Background(), "SELECT id, content, created_at FROM pastes ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var pastes []Paste
		for rows.Next() {
			var p Paste
			err := rows.Scan(&p.ID, &p.Content, &p.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			pastes = append(pastes, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pastes)
	} else if r.Method == "POST" {
		var p Paste
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow(context.Background(), "INSERT INTO pastes (content) VALUES ($1) RETURNING id, created_at", p.Content).Scan(&p.ID, &p.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
