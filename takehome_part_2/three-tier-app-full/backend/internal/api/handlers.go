package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "db not ok", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rows, err := db.Query(`SELECT id, title, done, created_at FROM todos ORDER BY id DESC`)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var list []Todo
			for rows.Next() {
				var t Todo
				if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				list = append(list, t)
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(list)

		case http.MethodPost:
			var payload struct {
				Title string `json:"title"`
			}
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				http.Error(w, "invalid json", http.StatusBadRequest)
				return
			}
			if payload.Title == "" {
				http.Error(w, "title is required", http.StatusBadRequest)
				return
			}
			var t Todo
			err := db.QueryRow(
				`INSERT INTO todos (title) VALUES ($1) RETURNING id, title, done, created_at`,
				payload.Title,
			).Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt)
			if err != nil {
				log.Println("insert error:", err)
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(t)

		default:
			w.Header().Set("Allow", "GET, POST, OPTIONS")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
