package model

import (
	"errors"
	"strings"
	"time"
)

// Publication posted by user
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      string    `json:"likes,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (publication *Publication) Prepare() error {
	if err := publication.valid(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publication) valid() error {
	if publication.Title == "" {
		return errors.New("the title is mandatory and cannot be blank")
	}

	if publication.Content == "" {
		return errors.New("the content is mandatory and cannot be blank")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
