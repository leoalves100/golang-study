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

func (repository Publication) SearchPublicationID(publicationID uint64) (model.Publication, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from
		publications p inner join users u
		on u.id = p.author_id where p.id = ?`,
		publicationID)

	if err != nil {
		return model.Publication{}, err
	}
	defer line.Close()

	var publication model.Publication

	if line.Next() {
		if err = line.Scan(
			&publication.ID,
			&publication.AuthorID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UpdatedAt,
			&publication.AuthorNick,
		); err != nil {
			return model.Publication{}, err
		}
	}

	return publication, nil
}

func (repository Publication) SearchPublications(userID uint64) ([]model.Publication, error) {
	line, err := repository.db.Query(`
	select p.*, u.nick from publications p 
	inner join users u on u.id = p.author_id
	inner join followers f on p.author_id = f.user_id
	where u.id = ? or f.followers_id = ?
	order by 1 desc`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer line.Close()

	var publications []model.Publication

	for line.Next() {
		var publication model.Publication

		if err = line.Scan(
			&publication.ID,
			&publication.AuthorID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UpdatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}
