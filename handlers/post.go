package handlers

import (
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

	added, err := db.GetProductById(prod.ID)
	if err != nil {
		http.Error(w, "unable to get product", http.StatusInternalServerError)
		return
	}

	if err := added.ToJson(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}
