package handlers

import (
	"context"
	"fmt"
	"net/http"

	"MicroEcom/db"
)

type KeyProduct struct{}

func (p *Products) MiddelwarePoroductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &db.Product{}

		if err := prod.FromJson(r.Body); err != nil {
			p.l.Println("[Error] reading product from request body", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		if err := prod.Validate(); err != nil {
			p.l.Println("[Error] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Invalid product data: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:8080" || origin == "http://127.0.0.1:8080" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
