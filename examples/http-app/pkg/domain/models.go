package domain

type (
	User struct {
		ID           int    `json:"id,omitempty"`
		Email        string `json:"email"`
		FirstName    string `json:"firstname"`
		LastName     string `json:"lastname"`
		HashPassword []byte `json:"hashpassword,omitempty"`
	}

	Product struct {
		ID             int     `json:"id"`
		SKU            string  `json:"sku"`
		Name           string  `json:"name"`
		DiscountPerc   float64 `json:"discountPerc"`
		DiscountAmount float64 `json:"discountAmount"`
	}

	// Define all struct types here
)
