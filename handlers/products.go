// # Documentation for MicroEcom
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
)

// A single product returned in the response
type Products struct {
	l *log.Logger
}

// NewProducts creates a new Products handler with the provided logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
