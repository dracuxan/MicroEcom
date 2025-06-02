package db

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Test Product",
		Price: 10.0,
		SKU:   "AB-01",
	}
	err := p.Validate()
	if err != nil {
		t.Errorf("Expected validation to pass, got error: %v", err)
	}
}
