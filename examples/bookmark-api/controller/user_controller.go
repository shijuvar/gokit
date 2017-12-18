package controller

import (
	"encoding/json"
	"net/http"

	utils "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
	"github.com/shijuvar/gokit/examples/bookmark-api/model"
	"github.com/shijuvar/gokit/examples/bookmark-api/store"
)

// Register add a new User document
// Handler for HTTP Post - "/users/register"
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	userModel := dataResource.Data
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	user := model.User{
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	}
	// Insert User document
	userStore.Create(user, userModel.Password)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginUser := dataResource.Data
	dataStore := NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("users")
	userStore := store.UserStore{C: col}
	// Authenticate the login user
	user, err := userStore.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}
	// Generate JWT token
	token, err = utils.GenerateJWT(user.Email, "member")
	if err != nil {
		utils.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
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
	w.Write(j)
}
