package domain

type (
	UserStore interface {
		Create(User, string) (User, error)
		Login(string, string) (User, error)
	}
	ProductStore interface {
		Create(Product) (Product, error)
		// ToDo: Define all contract methods here
	}
)
