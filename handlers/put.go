package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"MicroEcom/db"
)

// swagger:route PUT /products/{id} Products updateProduct
// # UpdateProduct updates a product by ID
// responses:
//	 200: updateProductResponse
//	 404: errorResponse
//	 500: errorResponse

// UpdateProduct updates a product by ID in the database
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

	if err := prod.ToJson(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}
