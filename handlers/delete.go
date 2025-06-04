package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"MicroEcom/db"
)

// swagger:route DELETE /products/{id} Products deleteProduct
// # DeleteProduct deletes a product by ID
// responses:
//	 200: deleteProductResponse
//	 404: errorResponse
//	 500: errorResponse

// DeleteProduct deletes a product by ID from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := db.DeleteProduct(id); err == db.ErrProductNotFound {
		p.l.Println("[ERROR] invalid id")
		http.Error(w, "Product not found", http.StatusNotFound)

		return
	}

	p.l.Println("Handle DELETE Product:", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": fmt.Sprintf("Product deleted successfully ID: %d", id),
	})
}
