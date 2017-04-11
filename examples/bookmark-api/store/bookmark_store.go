package store

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shijuvar/gokit/examples/bookmark-api/model"
)

// BookmarkStore provides CRUD operations against the collection "bookmarks".
type BookmarkStore struct {
	C *mgo.Collection
}

// Create inserts the value of struct Bookmark into collection.
func (store BookmarkStore) Create(b *model.Bookmark) error {
	// Assign a new bson.ObjectId
	b.ID = bson.NewObjectId()
	b.CreatedOn = time.Now()
	err := store.C.Insert(b)
	return err
}

// Update modifies an existing document of a collection.
func (store BookmarkStore) Update(b model.Bookmark) error {
	// partial update on MogoDB
	err := store.C.Update(bson.M{"_id": b.ID},
		bson.M{"$set": bson.M{
			"name":        b.Name,
			"description": b.Description,
			"location":    b.Location,
			"priority":    b.Priority,
			"tags":        b.Tags,
		}})
	return err
}

// Delete removes an existing document from the collection.
func (store BookmarkStore) Delete(id string) error {
	err := store.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// GetAll returns all documents from the collection.
func (store BookmarkStore) GetAll() []model.Bookmark {
	var b []model.Bookmark
	iter := store.C.Find(nil).Sort("priority", "-createdon").Iter()
	result := model.Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

// GetByUser returns all documents from the collection.
func (store BookmarkStore) GetByUser(user string) []model.Bookmark {
	var b []model.Bookmark
	iter := store.C.Find(bson.M{"createdby": user}).Sort("priority", "-createdon").Iter()
	result := model.Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

// GetByID returns a single document from the collection.
func (store BookmarkStore) GetByID(id string) (model.Bookmark, error) {
	var b model.Bookmark
	err := store.C.FindId(bson.ObjectIdHex(id)).One(&b)
	return b, err
}

// GetByTag returns all documents from the collection filtering by tags.
func (store BookmarkStore) GetByTag(tags []string) []model.Bookmark {
	var b []model.Bookmark
	iter := store.C.Find(bson.M{"tags": bson.M{"$in": tags}}).Sort("priority", "-createdon").Iter()
	result := model.Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}
