package handlers

import (
	"net/http"

	"MicroEcom/db"
)

// swagger:route GET /products products getProducts
// GetProducts handles the retrieval of products.
// responses:
//   200: productsResponse

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
