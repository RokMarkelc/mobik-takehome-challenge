package main

import (
	"log"
	"net/http"
	"os"

	"github.com/example/three-tier/internal/api"
	"github.com/example/three-tier/internal/db"
	"github.com/example/three-tier/internal/middleware"
)

func main() {
	// Connect DB
	conn, err := db.ConnectFromEnv()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close()

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, conn)

	// Wrap with CORS middleware
	handler := middleware.CORS(mux)

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}
	log.Printf("Go API listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
