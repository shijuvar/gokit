package main

import (
	"time"

	r "github.com/dancannon/gorethink"
)

// Bookmark type reperesents the metadata of a bookmark.
type Bookmark struct {
	ID                          string `gorethink:"id,omitempty" json:"id"`
	Name, Description, Location string
	Priority                    int // Priority (1 -5)
	CreatedOn                   time.Time
	Tags                        []string
}

// BookmarkStore provides CRUD operations against the Table "bookmarks".
type BookmarkStore struct {
	Session *r.Session
}

// Create inserts the value of struct Bookmark into Table.
func (store BookmarkStore) Create(b *Bookmark) error {

	resp, err := r.Table("bookmarks").Insert(b).RunWrite(store.Session)
	if err == nil {
		b.ID = resp.GeneratedKeys[0]
	}

	return err
}

// Update modifies an existing value of a Table.
func (store BookmarkStore) Update(b Bookmark) error {

	var data = map[string]interface{}{
		"description": b.Description,
		"location":    b.Location,
		"priority":    b.Priority,
		"tags":        b.Tags,
	}
	// partial update on RethinkDB
	_, err := r.Table("bookmarks").Get(b.ID).Update(data).RunWrite(store.Session)
	return err
}

// Delete removes an existing value from the Table.
func (store BookmarkStore) Delete(id string) error {
	_, err := r.Table("bookmarks").Get(id).Delete().RunWrite(store.Session)
	return err
}

// GetAll returns all documents from the Table.
func (store BookmarkStore) GetAll() ([]Bookmark, error) {
	bookmarks := []Bookmark{}

	res, err := r.Table("bookmarks").OrderBy("priority", r.Desc("createdon")).Run(store.Session)
	err = res.All(&bookmarks)
	return bookmarks, err
}

// GetByID returns single document from the Table.
func (store BookmarkStore) GetByID(id string) (Bookmark, error) {
	var b Bookmark
	res, err := r.Table("bookmarks").Get(id).Run(store.Session)
	res.One(&b)
	return b, err
}

// // GetByTag returns all documents from the collection filtering by tags.
// func (store BookmarkStore) GetByTag(tags []string) ([]Bookmark, error) {
// 	bookmarks := []Bookmark{}
// 	res, err := r.Table("bookmarks").Filter(func(row r.Term) r.Term {
// 		return r.Expr(tags).Contains(row.Field("tags"))
// 	}).Run(store.Session)
// 	err = res.All(&bookmarks)
// 	return bookmarks, err
// }
