package main

import (
	"fmt"
	"time"
)

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle, e.Location)
}

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

func main() {
	steve := Developer{
		Employee: Employee{
			FirstName: "Steve",
			LastName:  "John",
			Dob:       time.Date(1990, time.February, 17, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Software Engineer",
			Location:  "San Fancisco",
		},
		Skills: []string{"Go", "Docker", "Kubernetes"},
	}
	steve.PrintName()
	steve.PrintDetails()
}
