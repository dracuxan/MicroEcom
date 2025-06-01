package handlers

import (
	"log"
	"net/http"

	"MicroEcom/db"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter) {
	lp := db.GetProducts()

	if err := lp.ToJson(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}
