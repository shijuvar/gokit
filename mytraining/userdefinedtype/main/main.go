package main

import (
	"fmt"
	"niel-biswas/gokit/mytraining/userdefinedtype/domain"
	"niel-biswas/gokit/mytraining/userdefinedtype/mapstore"
)

// Organises the CRUD operations at UI layer
type CustomerController struct {
	// Explicit dependency that hides dependent logic
	store domain.CustomerStore // CustomerStore value
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created with ID", c.ID)
}

func (cc CustomerController) Update(k string, c domain.Customer) {
	err := cc.store.Update(k, c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("A Customer has been updated with ID", k)
}

func (cc CustomerController) Delete(k string) {
	err := cc.store.Delete(k)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("A Customer has been deleted with ID", k)
}

func (cc CustomerController) Fetch(k string) {
	c, err := cc.store.GetById(k)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(c)
}

func (cc CustomerController) FetchAll() {
	c, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, v := range c {
		fmt.Println(v)
	}
}

func main() {

	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
		// store : mongodb.NewMongoStore(), // with another database
	}
	var breakProgram bool = false

	var id, name, email, userInput string
	for breakProgram != true {
		fmt.Print("\nWelcome to ABC Company's customer data store. Please choose one of the options below :\n1. Create Customer (cc)                                                          \n2. Update Customer (uc)                                                          \n3. Delete Customer (dc)                                                          \n4. Get Customer (gc)                                                             \n5. Get All Customers (gac)                                                       \nPlease enter the short-hand notation (value inside brackets): ")
		fmt.Scanf("%s\n", &userInput)
		switch userInput {
		case "cc":
			fmt.Print("\nEnter Customer ID, Name and Email (seperated by a space): ")
			fmt.Scanf("%s %s %s\n", &id, &name, &email)
			customer := domain.Customer{
				ID:    id,
				Name:  name,
				Email: email,
			}
			controller.Add(customer) // Create new Customer
		case "uc":
			fmt.Print("\nEnter Customer ID, Name and Email for update (seperated by a space): ")
			fmt.Scanf("%s %s %s\n", &id, &name, &email)
			customer := domain.Customer{
				ID:    id,
				Name:  name,
				Email: email,
			}
			controller.Update(id, customer) // Update existing Customer
		case "dc":
			fmt.Print("\nEnter Customer ID for delete: ")
			fmt.Scanf("%s\n", &id)
			controller.Delete(id) // Delete existing Customer
		case "gc":
			fmt.Print("\nEnter Customer ID for fetching customer from the data store: ")
			fmt.Scanf("%s\n", &id)
			controller.Fetch(id) // Fetch existing Customer
		case "gac":
			controller.FetchAll() // Fetch all Customers
		default:
			fmt.Println("unknown user input, please try again.")
		}

		fmt.Print("\nDo you want to Continue or Quit? Enter c to Continue or q to Quit: ")
		fmt.Scanf("%s\n", &userInput)

		if userInput == "q" {
			breakProgram = true
		}
	}

}
