package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	utils "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
	"github.com/shijuvar/gokit/examples/bookmark-api/store"
)

// CreateBookmark insert a new Bookmark.
// Handler for HTTP Post - "/bookmarks
func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	var dataResource BookmarkResource
	// Decode the incoming Bookmark json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid Bookmark data",
			500,
		)
		return
	}
	bookmark := &dataResource.Data
	// Creates a new DatStore value to working with MongoDB store.
	dataStore := NewDataStore()
	// Add to the mgo.Session.Close()
	defer dataStore.Close()
	// Get the mgo.Collection for "bookmarks"
	col := dataStore.Collection("bookmarks")
	// Creates an instance of BookmarkStore
	bookmarkStore := store.BookmarkStore{C: col}
	// Takes user name from Context
	user := r.Context().Value("user")
	if user != nil {
		bookmark.CreatedBy = user.(string)
	}
	// Insert a bookmark document
	err = bookmarkStore.Create(bookmark)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid Bookmark data",
			500,
		)
		return
	}
	j, err := json.Marshal(BookmarkResource{Data: *bookmark})
	// If error is occured,
	// Send JSON response using helper function common.DisplayAppError
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Write the JSON data to the ResponseWriter
	w.Write(j)

}

// GetBookmarks returns all Bookmark documents
// Handler for HTTP Get - "/Bookmarks"
func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("bookmarks")
	bookmarkStore := store.BookmarkStore{C: col}
	bookmarks := bookmarkStore.GetAll()
	j, err := json.Marshal(BookmarksResource{Data: bookmarks})
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetBookmarkByID returns a single bookmark document by id
// Handler for HTTP Get - "/Bookmarks/{id}"
func GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("bookmarks")
	bookmarkStore := store.BookmarkStore{C: col}

	bookmark, err := bookmarkStore.GetByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)

		} else {
			utils.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)

		}
		return
	}
	j, err := json.Marshal(bookmark)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetBookmarksByUser returns all Bookmarks created by a User
// Handler for HTTP Get - "/Bookmarks/users/{id}"
func GetBookmarksByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("bookmarks")
	bookmarkStore := store.BookmarkStore{C: col}
	bookmarks := bookmarkStore.GetByUser(user)
	j, err := json.Marshal(BookmarksResource{Data: bookmarks})
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateBookmark update an existing Bookmark document
// Handler for HTTP Put - "/Bookmarks/{id}"
func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource BookmarkResource
	// Decode the incoming Bookmark json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid Bookmark data",
			500,
		)
		return
	}
	bookmark := dataResource.Data
	bookmark.ID = id
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("bookmarks")
	bookmarkStore := store.BookmarkStore{C: col}
	// Update an existing Bookmark document
	if err := bookmarkStore.Update(bookmark); err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// DeleteBookmark deletes an existing Bookmark document
// Handler for HTTP Delete - "/Bookmarks/{id}"
func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("bookmarks")
	bookmarkStore := store.BookmarkStore{C: col}
	// Delete an existing Bookmark document
	err := bookmarkStore.Delete(id)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
