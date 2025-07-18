package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/prathish-ks/openbank-connect/internal/handler"
)

func main() {
    r := chi.NewRouter()

    r.Get("/health", handler.HealthCheck)
	r.Get("/auth/start", handler.StartAuth)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
