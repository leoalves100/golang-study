package model

// User Representa um usuário
type User struct {
	ID       uint32 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
}
