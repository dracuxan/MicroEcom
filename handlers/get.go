package handlers

import (
	"net/http"

	"MicroEcom/db"
)

// swagger:route GET /products Products GetProducts
// GetProducts handles the retrieval of products.
// responses:
//   200: productsResponse
//   500: errorResponse
//   405: errorResponse

// GetProducts retrieves all products from the database and writes them to the response.
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p.l.Println("Handle GET Products")
	w.Header().Set("Content-Type", "application/json")
	lp := db.GetProducts()

	if err := lp.ToJson(w); err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}
