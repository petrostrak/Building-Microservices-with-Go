package handlers

import (
	"net/http"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(prod)
}
