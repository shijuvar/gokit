package main

import (
	"encoding/json"
	"fmt"
)

// Employee struct with struct tags
type Employee struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	JobTitle  string `json:"job"`
}

func main() {
	emp := Employee{
		FirstName: "Shiju",
		LastName:  "Varghese",
		JobTitle:  "Architect",
	}
	// Encoding to JSON
	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr := string(data)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	b := []byte(`{"id":101,"firstname":"Irene","lastname":"Rose","job":"Developer"}`)
	var emp1 Employee
	// Decoding JSON to a struct type
	err = json.Unmarshal(b, &emp1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("The Employee value is:")
	fmt.Printf("ID:%d, Name:%s %s, JobTitle:%s", emp1.ID, emp1.FirstName, emp1.LastName, emp1.JobTitle)
}
