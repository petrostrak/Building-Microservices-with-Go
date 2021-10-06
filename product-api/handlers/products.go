// Package classification of Product API
//
// Documentation for Product API
//
// Schemas: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger: meta

package handlers

import (
	"log"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Products is a http.Handler
type Product struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

type KeyProduct struct{}
