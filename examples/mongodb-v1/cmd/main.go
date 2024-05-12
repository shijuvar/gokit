package main

import (
	"fmt"

	"mongodb/model"
	"mongodb/repository"
)

func main() {
	repository, err := repository.NewNoteRepository()
	if err != nil {
		fmt.Println(err)
		return
	}
	note := model.Note{
		Title:       "example note",
		Description: "example note description",
	}
	if id, err := repository.Create(note); err == nil {
		fmt.Println("New record has been created with id:", id)
		qNote, _ := repository.GetById(id)
		fmt.Printf("Query Note:%+v\n", qNote)
	}
}
