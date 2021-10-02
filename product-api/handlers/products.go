package handlers

import (
	"log"
	"net/http"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/data"
)

// Products is a http.Handler
type Product struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	}

	// handle an update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "cannot encode to json", http.StatusInternalServerError)
	}
}
