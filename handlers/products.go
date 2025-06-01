package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 || len(g[0]) != 2 {
			p.l.Println("Invalid URL more than 1 id")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		ldString := g[0][1]
		id, err := strconv.Atoi(ldString)
		if err != nil {
			p.l.Println("Invalid URL unable to convert to number")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
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

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &db.Product{}
	if err := prod.FromJson(r.Body); err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

	db.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &db.Product{}
	if err := prod.FromJson(r.Body); err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}

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
