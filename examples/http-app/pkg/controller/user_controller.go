package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %w", err)
	}
	// Persistence
	newUser, err := handler.Store.Create(user.User, user.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Error on inserting User: %w", err)
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
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %w", err)

	}
	// Authenticate the login loginUser
	user, err := handler.Store.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("Authentication failed: %w", err)

	}
	// Generate JWT token if login is successful
	token, err = auth.GenerateJWT(user.Email, "member")
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Error on generating the token: %w", err)

	}
	authUser := authUserModel{User: user, Token: token}
	return authUser, http.StatusOK, nil

}
