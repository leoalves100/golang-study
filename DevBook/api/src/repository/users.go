package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {
	return &User{db}
}

// CreateUsers Create a user in the database
func (repository User) CreateUsers(user model.User) (uint64, error) {
	// Creates a statement to be executed
	statment, err := repository.db.Prepare("insert into users (name, nick, mail, password) values (?, ?, ?, ?)")
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

	return uint64(lastInsertID), nil
}

// SearchUsers Search 1 or n* user in the database, based on nike or name
func (repository User) SearchUsers(nameORnick string) ([]model.User, error) {
	//Create a string with percentage in prefix and suffix. (Ex: '%name%')
	nameORnick = fmt.Sprintf("%%%s%%", nameORnick)

	lines, err := repository.db.Query(
		"select id, name, nick, mail from users "+
			"where name LIKE ? or nick LIKE ?", nameORnick, nameORnick,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []model.User

	for lines.Next() {
		var user model.User

		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Mail); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchUser Search 1 user in the database, based on nike or name
func (repository User) SearchUser(ID uint64) (model.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, mail, created_at, updated_at from users where id = ?", ID,
	)

	if err != nil {
		return model.User{}, err
	}

	defer lines.Close()

	var user model.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Mail,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// SearchUser Search 1 user specific mail in the database
func (repository User) SearchUserMail(Mail string) (model.User, error) {
	lines, err := repository.db.Query(
		"select id, mail, password from users where mail = ?", Mail,
	)

	if err != nil {
		return model.User{}, err
	}

	defer lines.Close()

	var user model.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Mail, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// UpdateUser update a specific user in the database
func (repository User) UpdateUser(ID uint64, user model.User) error {
	stat, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, mail = ? where id = ?",
	)
	if err != nil {
		return err
	}

	defer stat.Close()

	if _, err := stat.Exec(user.Name, user.Nick, user.Mail, ID); err != nil {
		return err
	}

	return nil
}

// DeleteUser delete a specific user in the database
func (repository User) DeleteUser(ID uint64) error {
	stat, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}

	defer stat.Close()

	if _, err := stat.Exec(ID); err != nil {
		return err
	}

	return nil
}

// Follow is a function that allows a user to follow another user.
//
// It takes two parameters, userID and followID, both of type uint64.
// It returns an error if there was an issue while following the user.
// Ref: https://www.mysqltutorial.org/mysql-insert-ignore/
func (repository User) Follow(userID, followID uint64) error {
	statement, err := repository.db.Prepare("insert ignore into followers (user_id, followers_id) values (?, ?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

func (repository User) UnFollow(userID, followID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and followers_id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

func (repository User) SearchFollowers(userID uint64) ([]model.User, error) {
	statement, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.mail, u.created_at from users u 
		inner join followers s on u.id = s.followers_id
		where s.user_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var users []model.User

	for statement.Next() {
		var user model.User

		if err = statement.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Mail,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository User) SearchFollowing(userID uint64) ([]model.User, error) {
	statement, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.mail, u.created_at from users u 
		inner join followers s on u.id = s.user_id
		where s.followers_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var users []model.User

	for statement.Next() {
		var user model.User

		if err = statement.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Mail,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
