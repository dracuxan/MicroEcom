package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"MicroEcom/db"
)

// swagger:route POST /products Products createProduct
// CreateProduct handles the creation of a new products.
// responses:
//	201: newProductResponse
//	400: errorResponse
//	500: errorResponse

// Create handles the creation of a new product.
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*db.Product)

	db.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": fmt.Sprintf("Product created successfully ID: %d", prod.ID),
	})
}
