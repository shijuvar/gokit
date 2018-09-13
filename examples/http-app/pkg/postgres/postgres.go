package postgres

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Config holds the configuration used for instantiating a new DataStore.
type Config struct {
	Host, Port, User, Password, Database string
}

type DataStore struct {
	Db  *sql.DB
	cfg Config
}

// New returns a DataStore instance with the sql.DB set with the postgres
func New(cfg Config) (DataStore, error) {
	var err error
	var store DataStore

	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err = errors.Errorf(
			"All fields must be set (%s)",
			spew.Sdump(cfg))
		return store, err
	}
	store.cfg = cfg
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))

	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to postgre database (%s)",
			spew.Sdump(cfg)) // Sdump returns a string with the passed arguments formatted exactly the same as Dump.
		return store, err
	}

	// Ping verifies if the connection to the database is alive or if a
	// new connection can be made.
	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err,
			"Couldn't ping postgre database (%s)",
			spew.Sdump(cfg))
		return store, err
	}

	store.Db = db
	return store, nil
}
