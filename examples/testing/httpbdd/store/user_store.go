package store

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/model"
)

// MongoDB Session
var mgoSession *mgo.Session

// Create a MongoDB Session
func createDBSession() {
	var err error
	mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:   []string{"127.0.0.1"},
		Timeout: 60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

// Initializes the MongoDB Session
func init() {
	createDBSession()
}

// MongoUserStore provides persistence logic for "users" collection.
type MongoUserStore struct{}

// AddUser insert new User
func (store *MongoUserStore) AddUser(user model.User) error {
	session := mgoSession.Copy()
	defer session.Close()
	userCol := session.DB("userdb").C("users")
	// Check whether email id is exists or not
	var existUser model.User
	err := userCol.Find(bson.M{"email": user.Email}).One(&existUser)
	if err != nil {
		if err == mgo.ErrNotFound { // Email is unique, no records found
		}
	}
	if (model.User{}) != existUser { // there is a user
		return model.ErrorEmailExists
	}
	err = userCol.Insert(user)
	return err
}

// GetUsers returns all documents from the collection.
func (store *MongoUserStore) GetUsers() []model.User {
	session := mgoSession.Copy()
	defer session.Close()
	userCol := session.DB("userdb").C("users")
	var users []model.User
	iter := userCol.Find(nil).Iter()
	result := model.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}
