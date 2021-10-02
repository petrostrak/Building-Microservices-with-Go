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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}
	// handle an update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	if err := prod.FromJSON(r.Body); err != nil {
		http.Error(rw, "cannot decode from json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

// getProducts returns the products from the data store
func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "cannot encode to json", http.StatusInternalServerError)
	}
}
