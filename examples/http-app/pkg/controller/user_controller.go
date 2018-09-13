package controller

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/shijuvar/gokit/examples/http-app/pkg/auth"
	"github.com/shijuvar/gokit/examples/http-app/pkg/domain"
)

type UserController struct {
	Store domain.UserStore // This gives extensibility and testability
}

// HTTP Post - /users
func (handler UserController) PostUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var user createUser
	// Decode the incoming json data to note struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(err, "Unable to decode JSON request body")
	}
	// Persistence
	newUser, err := handler.Store.Create(user.User, user.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Error on inserting User")
	}
	return newUser, http.StatusCreated, nil
}

// Login authenticates the HTTP request with username and password
// Handler for HTTP Post - "/users/login"
func (handler UserController) PostLogin(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var loginUser loginUser
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		return nil, http.StatusBadRequest, errors.Wrap(err, "Unable to decode JSON request body")

	}
	// Authenticate the login loginUser
	user, err := handler.Store.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.Wrap(err, "Authentication failed")

	}
	// Generate JWT token if login is successful
	token, err = auth.GenerateJWT(user.Email, "member")
	if err != nil {
		return nil, http.StatusInternalServerError, errors.Wrap(err, "Error on generating the token")

	}
	authUser := authUserModel{User: user, Token: token}
	return authUser, http.StatusOK, nil

}
