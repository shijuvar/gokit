package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const dtotag = "dto"

var TYPE_REGISTRY = make(map[string]reflect.Type)

type Person struct {
	ID          int    `dto:"ID"`
	Name        string `dto:"Name"`
	Phone       string
	DateOfBirth time.Time `dto:"DoB"`
}

type PersonDTO struct {
	ID          int       `dto:"ID" json:"ID"`
	Name        string    `dto:"Name" json:"Name"`
	DateOfBirth time.Time `dto:"DoB"`
}

func ToDTO(model interface{}) interface{} {

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get(dtotag)
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	val := reflect.New(TYPE_REGISTRY[reflect.TypeOf(model).String()+"DTO"]).Elem()

	// find val's fields
	for i := 0; i < val.NumField(); i++ {
		fieldInfo := val.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get(dtotag)
		for k, v := range fields {
			if k == name {
				val.Field(i).Set(v)
				break
			}
		}
	}
	return val.Interface()
}

func (p Person) ToDTO() PersonDTO {
	dto := PersonDTO{ID: p.ID, Name: p.Name, DateOfBirth: p.DateOfBirth}
	return dto
}

func main() {
	TYPE_REGISTRY[reflect.TypeOf(Person{}).String()] = reflect.TypeOf(Person{})
	TYPE_REGISTRY[reflect.TypeOf(PersonDTO{}).String()] = reflect.TypeOf(PersonDTO{})

	model := Person{ID: 7, Name: "James Bond", Phone: "123-456-7890", DateOfBirth: time.Now()}
	val := ToDTO(model)

	fmt.Println(model)

	// cast val to DTO
	dto := val.(PersonDTO)
	fmt.Println(dto)

	checkdto := model.ToDTO()

	fmt.Println(reflect.DeepEqual(dto, checkdto))
}
