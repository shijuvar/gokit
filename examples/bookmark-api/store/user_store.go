package store

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shijuvar/gokit/examples/bookmark-api/model"
)

// UserStore provides persistence logic for "users" collection.
type UserStore struct {
	C *mgo.Collection
}

// Create insert new User
func (store UserStore) Create(user model.User, password string) error {

	user.ID = bson.NewObjectId()
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = hpass
	err = store.C.Insert(user)
	return err
}

// Login authenticates the User
func (store UserStore) Login(email, password string) (model.User, error) {
	var user model.User
	err := store.C.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return model.User{}, err
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
