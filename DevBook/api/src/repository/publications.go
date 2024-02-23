package repository

import (
	"api/src/model"
	"database/sql"
)

type Publication struct {
	db *sql.DB
}

func NewRepositoryPublications(db *sql.DB) *Publication {
	return &Publication{db}
}

func (repository Publication) CreatePublication(publication model.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into publications (title, content, author_id) values (?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	latestIDInsert, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(latestIDInsert), nil
}
