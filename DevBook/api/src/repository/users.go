package repository

import (
	"api/src/model"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) CreateUsers(user model.User) (uint32, error) {
	// Creates a statement to be executed
	statment, err := repository.db.Prepare("insert into usuarios (name, nick, mail, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statment.Close()

	// Executes statement by passing the user values
	res, err := statment.Exec(user.Name, user.Nick, user.Mail, user.Password)
	if err != nil {
		return 0, err
	}

	// Get last ID user
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(lastInsertID), nil
}
