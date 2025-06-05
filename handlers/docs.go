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

// Payload for creating a new product
// swagger:parameters createProduct
type newProductParameterWrapper struct {
	// The product to create
	// in: body
	// required: true
	Product db.Product `json:"product"`
}

// ID of the product to update
// swagger:parameters updateProduct
type productIDUpdateParameterWrapper struct {
	// The ID of the product to update
	// in: path
	// required: true
	ID int `json:"id"`
}

// Payload for updating an existing product
// swagger:parameters updateProduct
type updatedProductParameterWrapper struct {
	// The product to update
	// in: body
	// required: true
	Product db.Product `json:"product"`
}

// Response for the delete operation
// swagger:response deleteProductResponse
type deleteProductResponseWrapper struct {
	// A message indicating the result of the delete operation
	// in: body
	Body GenericMessage
}

type GenericMessage struct {
	Message string `json:"message"`
}

// Error response for various operations
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Updated product response with the product ID
// swagger:response updateProductResponse
type updateProductResponseWrapper struct {
	// Message with the updated product ID
	// in: body
	Body GenericMessage
}

// New product response with the product ID
// swagger:response newProductResponse
type newProductResponseWrapper struct {
	// The newly created product
	// in: body
	Body GenericMessage
}
