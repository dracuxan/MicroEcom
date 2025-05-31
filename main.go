package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "fucl", http.StatusBadRequest)
		return
	}

	log.Printf("Data received: %s\n", d)
	fmt.Fprintf(w, "Hello %s!\n", d)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":9090", nil)
}
