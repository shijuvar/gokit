package controller

import "github.com/shijuvar/gokit/examples/http-app/pkg/domain"

// Data models used for API endpoints
type (
	// userResource For Post - /users/register
	createUser struct {
		User     domain.User `json:"user"`
		Password string      `json:"password"`
	}
	// loginUser for authentication
	loginUser struct {
		Email    string `json:"user_id"`
		Password string `json:"password"`
	}
	// authUserModel for authorized user with access token
	authUserModel struct {
		User  domain.User `json:"user"`
		Token string      `json:"token"`
	}
)
