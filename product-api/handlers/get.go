package handlers

import (
	"net/http"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//
// GetProducts returns the products from the data store
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	if err := lp.ToJSON(rw); err != nil {
		http.Error(rw, "cannot encode to json", http.StatusInternalServerError)
	}
}
