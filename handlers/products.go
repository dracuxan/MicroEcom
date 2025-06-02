package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"MicroEcom/db"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p.l.Println("Handle GET Products")
	lp := db.GetProducts()

	if err := lp.ToJson(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*db.Product)

	db.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, er := strconv.Atoi(vars["id"])

	if er != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(*db.Product)

	p.l.Println("Handle PUT Product for id:", id)

	err := db.UpdateProduct(id, prod)
	if err == db.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Updated Product: %#v", prod)
}

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
