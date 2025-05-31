package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Greet handles the HTTP request and responds with a greeting message.
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello handler with the provided logger.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the http.Handler interface for Hello.
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "fucl", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s!\n", d)
}
