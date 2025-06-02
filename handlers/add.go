package handlers

import (
	"net/http"

	"MicroEcom/db"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*db.Product)

	db.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}
