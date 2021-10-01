package handlers

import (
	"log"
	"net/http"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	}

	// handle an update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "cannot encode to json", http.StatusInternalServerError)
	}
}
