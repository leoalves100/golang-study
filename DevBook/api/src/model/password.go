package model

// Password template for user password update
type Password struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
