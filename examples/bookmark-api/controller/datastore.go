package controller

import (
	"gopkg.in/mgo.v2"

	"github.com/shijuvar/gokit/examples/bookmark-api/bootstrapper"
)

// DataStore for MongoDB
type DataStore struct {
	MongoSession *mgo.Session
}

// Close closes a mgo.Session value.
// Used to add defer statements for closing the copied session.
func (ds *DataStore) Close() {
	ds.MongoSession.Close()
}

// Collection returns mgo.collection for the given name
func (ds *DataStore) Collection(name string) *mgo.Collection {
	return ds.MongoSession.DB(bootstrapper.AppConfig.Database).C(name)
}

// NewDataStore creates a new DataStore object to be used for each HTTP request.
func NewDataStore() *DataStore {
	session := bootstrapper.GetSession().Copy()
	dataStore := &DataStore{
		MongoSession: session,
	}
	return dataStore
}
