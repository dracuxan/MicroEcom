package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"MicroEcom/db"
)

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
