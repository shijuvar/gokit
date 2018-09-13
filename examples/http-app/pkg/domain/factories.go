package domain

// NewProduct returns an instance of Product by providing default values
func NewProduct() Product {
	var product Product
	// Provide all default values
	product.DiscountPerc = 0.00
	product.DiscountAmount = 0.00
	return product
}
