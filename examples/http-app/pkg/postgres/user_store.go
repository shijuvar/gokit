package postgres

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"database/sql"
	"log"

	"github.com/shijuvar/gokit/examples/http-app/pkg/domain"
)

// UserStore provides persistence logic for "users" table
type UserStore struct {
	Store DataStore
}

// Create creates a new User
func (userStore UserStore) Create(user domain.User, password string) (domain.User, error) {

	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return user, errors.Wrap(err, "Error on hashing password")
	}
	user.HashPassword = hpass
	sqlStatement := `
		INSERT INTO users (email_id, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id;`
	id := 0
	// Use Db.Exec if you don't want returning ID
	err = userStore.Store.Db.QueryRow(sqlStatement, user.Email, user.FirstName, user.LastName, user.HashPassword).Scan(&id)
	if err != nil {
		user.ID = id // assign returning ID
		return user, errors.Wrap(err, "Error while inserting on users")
	}
	return user, nil
}

// Login authenticates the User
func (userStore UserStore) Login(email, password string) (domain.User, error) {
	var user domain.User
	var err error
	sqlStatement := `SELECT first_name,last_name,password_hash FROM users where email_id=$1;`
	log.Printf(email)
	row := userStore.Store.Db.QueryRow(sqlStatement, email)

	switch err = row.Scan(&user.FirstName, &user.LastName, &user.HashPassword); err {
	case sql.ErrNoRows:
		err = errors.Wrap(err, "Invalid Email Id")
	case nil:
		// Validate password
		err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(password))
		if err != nil {
			err = errors.Wrap(err, "Invalid Password")
		}
	default:
		err = errors.Wrap(err, "Error on querying data")
	}

	return user, err
}
