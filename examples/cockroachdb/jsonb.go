package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type profile struct {
	ProfileID string
	UpdatedOn time.Time
	user
}
type user struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Location  string `json:"location,omitempty"`
	Status    string `json:"status,omitempty"`
}

func main() {
	// Connect to the "testdb" database.
	db, err := sql.Open("postgres", "postgresql://shijuvar@localhost:26257/testdb?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Create the "users" table.
	sql := "CREATE TABLE IF NOT EXISTS users (profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),last_updated TIMESTAMP DEFAULT now(),user_profile JSONB)"
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
	u := user{
		FirstName: "Shiju",
		LastName:  "Varghese",
		Location:  "Kochi",
		Status:    "Working",
	}
	// insert user
	err = insert(db, u)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New user has been created")
	// get all users
	profiles, err := get(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range profiles {
		log.Println("Profile ID:", v.ProfileID)
		log.Println("Updated On:", v.UpdatedOn.Format("02-Jan-2006 15:4:5"))
		log.Printf("User Profile: %v\n", v.user)
	}
}

func insert(db *sql.DB, u user) error {
	json, _ := json.Marshal(u)
	strJson := string(json)
	sql := "INSERT INTO users (user_profile) VALUES ($1)"
	if _, err := db.Exec(sql, strJson); err != nil {
		return err
	}
	return nil
}

func get(db *sql.DB) ([]profile, error) {
	sql := "SELECT profile_id,last_updated," +
		"user_profile->'first_name',user_profile->'last_name',user_profile->'location',user_profile->'status' " +
		"FROM users where user_profile is not null"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var profiles []profile
	for rows.Next() {
		var p profile
		if err := rows.Scan(&p.ProfileID, &p.UpdatedOn, &p.user.FirstName,
			&p.user.LastName, &p.user.Location, &p.user.Status); err != nil {
			return profiles, err
		}
		profiles = append(profiles, p)
	}
	if err = rows.Err(); err != nil {
		return profiles, err
	}
	return profiles, nil
}
