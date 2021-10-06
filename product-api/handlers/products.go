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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
