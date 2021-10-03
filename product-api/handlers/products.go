package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	if err := prod.FromJSON(r.Body); err != nil {
		http.Error(rw, "cannot decode from json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
		return
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	if err := prod.FromJSON(r.Body); err != nil {
		http.Error(rw, "cannot decode from json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

// getProducts returns the products from the data store
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "cannot encode to json", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

// MiddlewareProductValidation validates the product in the request and calls next if ok
func (p *Product) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		if err := prod.FromJSON(r.Body); err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "error reading product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
