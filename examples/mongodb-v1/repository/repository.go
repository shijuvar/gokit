package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mongodb/model"
)

var ctx = context.TODO()
var mongoClient *mongo.Client

type NoteRepository struct {
	noteCollection *mongo.Collection
}

func NewNoteRepository() (model.Repository, error) {
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
	return &NoteRepository{noteCollection: collection}, nil
}

func (m *NoteRepository) Create(n model.Note) (string, error) {
	if en, _ := getNoteByTitle(m.noteCollection, n.Title); en.Title == n.Title {
		return "", model.ErrNoteExists
	}
	n.CreatedOn = time.Now()
	n.ID = primitive.NewObjectID()
	result, err := m.noteCollection.InsertOne(ctx, n)
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func getNoteByTitle(c *mongo.Collection, title string) (model.Note, error) {
	filter := bson.D{{"title", title}}
	var n model.Note
	err := c.FindOne(ctx, filter).Decode(&n)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return n, model.ErrNotFound
		}
		return n, err
	}

	return n, nil
}

func (m *NoteRepository) Update(id string, n model.Note) error {
	if _, err := m.GetById(id); errors.Is(err, model.ErrNotFound) {
		return err
	}
	objID, _ := getObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set",
		bson.D{
			{"title", n.Title},
			{"description", n.Description},
		},
	}}
	_, err := m.noteCollection.UpdateOne(ctx, filter, update)
	return err
}
func (m *NoteRepository) Delete(id string) error {
	if _, err := m.GetById(id); errors.Is(err, model.ErrNotFound) {
		return err
	}
	objID, _ := getObjectIDFromHex(id)

	filter := bson.D{{"_id", objID}}
	_, err := m.noteCollection.DeleteOne(ctx, filter)
	return err
}

func getObjectIDFromHex(hex string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(hex)
	return objID, err
}
func (m *NoteRepository) GetById(id string) (model.Note, error) {
	objID, _ := getObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}
	var n model.Note
	err := m.noteCollection.FindOne(ctx, filter).Decode(&n)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return n, model.ErrNotFound
		}
		return n, err
	}

	return n, nil
}

func (m *NoteRepository) GetAll() ([]model.Note, error) {
	filter := bson.D{{}}
	var notes []model.Note

	cur, err := m.noteCollection.Find(ctx, filter)
	if err != nil {
		return notes, err
	}

	for cur.Next(ctx) {
		var n model.Note
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
		return notes, model.ErrNotFound
	}

	return notes, nil
}
