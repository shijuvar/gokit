package mongodb

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shijuvar/gokit/examples/http/restapi/model"
)

var ctx = context.TODO()
var mongoClient *mongo.Client

type MongoNoteRepository struct {
	noteCollection *mongo.Collection
}

func NewMongoNoteRepository() (model.Repository, error) {
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
	return &MongoNoteRepository{noteCollection: collection}, nil
}

func (m *MongoNoteRepository) Create(n model.Note) error {
	// Create a Version 4 UUID.
	uid, _ := uuid.NewV4()
	n.NoteID = uid.String()
	n.CreatedOn = time.Now()
	n.ID = primitive.NewObjectID()
	_, err := m.noteCollection.InsertOne(ctx, n)
	return err
}

func (m *MongoNoteRepository) Update(id string, n model.Note) error {
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
func (m *MongoNoteRepository) Delete(id string) error {
	filter := bson.D{{"noteid", id}}
	_, err := m.noteCollection.DeleteOne(ctx, filter)
	return err
}

func (m *MongoNoteRepository) GetById(id string) (model.Note, error) {
	filter := bson.D{{"noteid", id}}
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
func (m *MongoNoteRepository) GetAll() ([]model.Note, error) {
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
