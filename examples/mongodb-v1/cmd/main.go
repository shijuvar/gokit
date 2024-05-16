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

	// create, update and get
	if id, err := repository.Create(note); err == nil {
		fmt.Println("New record has been created with id:", id)
		if qNote, err := repository.GetById(id); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Query Note:%+v\n", qNote)
		}
	} else {
		fmt.Println("error while creating new Note:", err)
	}
	// trying to create with duplicate Note
	if _, err := repository.Create(note); err != nil {
		fmt.Println("error while creating new Note:", err)
	}

	note1 := model.Note{
		Title:       "mongodb",
		Description: "mongodb is a NoSQL document database",
	}
	id, err := repository.Create(note1)
	if err == nil {
		fmt.Println("New record has been created with id:", id)
	}
	// update
	note1.Title = "mongodb database"
	if err := repository.Update(id, note1); err != nil {
		fmt.Println("error while updating Note:", err)
	} else {
		fmt.Println("Updated Note")
	}
	// get all
	if notes, err := repository.GetAll(); err == nil {
		fmt.Println("printing notes:")
		for _, note := range notes {
			fmt.Printf("%+v\n", note)
		}
	}

	if err := repository.Delete(id); err != nil {
		fmt.Println("error while deleting Note:", err)
	} else {
		fmt.Println("Deleted Note")
	}
	// get all
	if notes, err := repository.GetAll(); err == nil {
		fmt.Println("printing notes after deletion:")
		for _, note := range notes {
			fmt.Printf("%+v\n", note)
		}
	}
}
