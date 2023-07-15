package model

import "time"

// Publication posted by user
type Publication struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Author    uint64    `json:"author,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
