package main

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var mongoClient *mongo.Client

type mongoNoteRepository struct {
	noteCollection *mongo.Collection
}

func newMongoNoteRepository() (repository, error) {
	if mongoClient == nil {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
		var err error
		mongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			return nil, err
		}
		err = mongoClient.Ping(ctx, nil)
		if err != nil {
			return nil, err
		}
	}
	collection := mongoClient.Database("notesdb").Collection("notes")
	return &mongoNoteRepository{noteCollection: collection}, nil
}

func (m *mongoNoteRepository) create(n note) error {
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.NoteID = uid.String()
	n.CreatedOn = time.Now()
	n.ID = primitive.NewObjectID()
	_, err := m.noteCollection.InsertOne(ctx, n)
	return err
}

func (m *mongoNoteRepository) update(id string, n note) error {
	filter := bson.D{{"noteid", id}}
	update := bson.D{{"$set",
		bson.D{
			{"title", n.Title},
			{"description", n.Description},
		},
	}}
	_, err := m.noteCollection.UpdateOne(ctx, filter, update)
	return err
}
func (m *mongoNoteRepository) delete(id string) error {
	filter := bson.D{{"noteid", id}}
	_, err := m.noteCollection.DeleteOne(ctx, filter)
	return err
}

func (m *mongoNoteRepository) getById(id string) (note, error) {
	filter := bson.D{{"noteid", id}}
	var n note
	err := m.noteCollection.FindOne(ctx, filter).Decode(&n)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return n, errNotFound
		}
		return n, err
	}

	return n, nil
}
func (m *mongoNoteRepository) getAll() ([]note, error) {
	filter := bson.D{{}}
	var notes []note

	cur, err := m.noteCollection.Find(ctx, filter)
	if err != nil {
		return notes, err
	}

	for cur.Next(ctx) {
		var n note
		err := cur.Decode(&n)
		if err != nil {
			return notes, err
		}

		notes = append(notes, n)
	}

	if err := cur.Err(); err != nil {
		return notes, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(notes) == 0 {
		return notes, errNotFound
	}

	return notes, nil
}
