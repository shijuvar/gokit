# Assignments: Professional Go

## Assignment 1: Basics and Functions
### Objective: To get familiarise Go basics

Do implement a simple calculator as being called in the main function:

```go
func main() {
    var result float64
    result = doCalculate(100, "+") // result => 100
    fmt.Println(result)
    result = doCalculate(50, "-") // result => 50
    fmt.Println(result)
    result = doCalculate(20, "/") // result => 2.5
    fmt.Println(result)
    result = doCalculate(10, "*") // result => 25
    fmt.Println(result)
}
```

## Assignment 2: Working with Collections
### Objective: To get familiarise Go basics with Map data structure

* Define a package level variable of type: 
```go 
   map[string]string
   ```
* Create functions for making insert, update, delete and get items to
and from the map (package level variable of type map) with the following
signature:

```go
addItem (k,v string)
updateItem (k,v string)
getById (k string)
getAll()
deleteItem (k string)
```

```go
// Declare package level variable for storing data in map
var data map[string]string
// init function will be automatically invoked before main function
// init function is used to initialise package level variables
func init() {
    data = make(map[string]string) // Initialise map with make
}
func addItem(k,v string) {
// ToDo: Check if key exists
    data [k] = v
}
```

## Assignment 3: Writing Idiomatic Go Code With User-Defined Type System
### Objective: Write idiomatic Go code with packages, struct and interface

#### Principles:
* SOLID and Clean Architecture
* Explicit Dependencies: Methods and classes should explicitly
require (typically through method parameters or constructor
parameters) any collaborating objects they need in order to
function correctly.
* Declarative Composition: Removes the dependent logic from
the composition process.
* "Be conservative in what you send, be liberal in what you
accept" — Robustness principle
* "Accept interfaces, return structs" –– A Go Proverb

**Create package named domain**
* Create a struct named Customer. 
* Create an interface named CustomerRepository to specify behaviours for CRUD on Customer.

```go
package domain 

type Customer struct {
    ID string
    Name string
    Email string
}
type CustomerRepository interface {
    Create (Customer) error
    Update (string, Customer) error
    Delete(string) error
    GetById(string) (Customer, error)
    GetAll() ([]Customer, error)
}
```

**Create package named memstore**
* Implement interface domain.CustomerRepository to persist domain.Customer data into a ```map[string]Customer```

```go
package memstore

import (
	"customerapp/domain"
)

// CustomerRepository provides memory based local data repository.
type CustomerRepository struct {
	repository map[string]domain.Customer
}

// NewCustomerRepository initialises the memory data repository
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{repository: make(map[string]domain.Customer)}
}

```
* Implement interface methods of domain.CustomerRepository

**Create package main**

* Create a CustomerController struct to organise the CRUD operations at UI layer
```go
// CustomerController Organises the CRUD operations at UI layer.
type CustomerController struct {
    repository domain.CustomerRepository
}

// Add function to add new customer.
func (cc CustomerController) Add(c domain.Customer) {
	err := cc.repository.Create(c)
	if err != nil {
		fmt.Println("Error: ", err)
	return
  }
  fmt.Println("New Customer has been created\n")
}
// Implement other methods for the type CustomerController

// Inside the main function
func main() {
    controller := CustomerController{ // initialize customer controller
        repository: memstore.NewCustomerRepository(),
        //repository: mongodb.NewCustomerRepository(), // switching to another persistent store
    }
    customer1 := domain.Customer{
        ID:    "cust101",
        Name:  "Rahul",
        Email: "rahul@gmail.com",
    }
controller.Add(customer1)

// Make all CRUD operations	
```

▪ By using CustomerController, make CRUD operations on Customer into an in-memory store.

## Assignment 4: HTTP Programming
### Objective: Write RESTful APIs  

**From assignment 3:**
* Add struct tag for JSON encoding package in the Customer entity of domain package.
* Create a package named controller and move the CustomerController type into this package in order to implement
http handlers as shown in the function signature below:
```go
func (ctl CustomerController) Post(w http.ResponseWriter, r *http.Request))
```
* Create a **package main** and start a new http server from it after configuring the routes with http.ServeMux.

## Assignment 5: TDD Unit Tests
### Objective: Write TDD unit tests
* Write TDD style unit tests for assignment 4

## Assignment 6: JWT Authentication to HTTP App
### Objective: Add JWT (https://jwt.io) authentication to HTTP REST API server and write HTTP middleware

**From assignment 4:**
* Create a a domain entity User to represent an authenticated user
* Write handler functions for register a new user and login to the
systems using user name and password
* When a user authenticates to the system using user name and
password, send back a JWT token to the HTTP client application
* Write a JWT based authorization middleware using Go JWT package
https://github.com/golang-jwt/jwt to restrict users to make CRUD
operations on Customer entity. In order to make CRUD operations on
Customer entity, a user should provide a valid JWT token (JWT
received after login to the system using user name and password) to
access the REST API endpoints of Customer entity.
* Write TDD styled unit tests for new functionalities.
