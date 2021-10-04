package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Petros",
		Price: 1.00,
		SKU:   "abc-abcd-abcdf",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
