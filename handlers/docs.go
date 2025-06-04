package handlers

import "MicroEcom/db"

// GenericError is a generic error message returned in the response
type GenericError struct {
	Message string `json:"message"`
}

// A list of products returned in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []db.Product
}

// ID of the product to delete
// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to delete
	// in: path
	// required: true
	ID int `json:"id"`
}

// Response for the delete operation
// swagger:response deleteProductResponse
type deleteProductResponseWrapper struct {
	// A message indicating the result of the delete operation
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Error response for various operations
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Updated product response
// swagger:response updateProductResponse
type updateProductResponseWrapper struct {
	// The updated product
	// in: body
	Body db.Product
}

// New product response
// swagger:response newProductResponse
type newProductResponseWrapper struct {
	// The newly created product
	// in: body
	Body db.Product
}
