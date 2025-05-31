package handlers

import (
	"log"
	"net/http"
)

// Bye handles the HTTP request and responds with a goodbye message.
type Bye struct {
	l *log.Logger
}

// NewBye creates a new Bye handler with the provided logger.
func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

// ServeHTTP implements the http.Handler interface for Bye.
func (b *Bye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.l.Println("sending bye")

	w.Write([]byte("Bye Bye\n"))
}
