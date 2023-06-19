package models

import (
	"log"
	"time"

	"github.com/ritvikdayal/todoGoWeb/db"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) Signup(userPayload User) (*User, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}
	// Prepare the SQL statement for inserting the user
	stmt, err := db.Prepare(`
		INSERT INTO users (id, name, birthday, gender, photo_url, current_time, active, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	// Execute the prepared statement with the user's data
	_, err = stmt.Exec(
		user.ID,
		user.Name,
		user.BirthDay,
		user.Gender,
		user.PhotoURL,
		user.Time,
		user.Active,
		user.UpdatedAt,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	stmt, err := db.Prepare(`
		SELECT id, name, birthday, gender, photo_url, current_time, active, updated_at
		FROM users
		WHERE id = ?
	`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	user := &User{}
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.BirthDay,
		&user.Gender,
		&user.PhotoURL,
		&user.Time,
		&user.Active,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
