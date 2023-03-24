package model

import (
	"api/src/sec"
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

// User Representa um usuário
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Mail      string `json:"mail,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.valid(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}

	return nil
}

func (user *User) valid(stage string) error {

	if user.Name == "" {
		return errors.New("the name field is mandatory")
	}

	if user.Nick == "" {
		return errors.New("the nick field is mandatory")
	}

	if user.Mail == "" {
		return errors.New("the mail field is mandatory")
	}

	if err := checkmail.ValidateFormat(user.Mail); err != nil {
		return fmt.Errorf("%v mail", err)
	}

	if stage == "register" && user.Password == "" {
		return errors.New("the password field is mandatory")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Mail = strings.TrimSpace(user.Mail)

	if stage == "register" {
		passwordHash, err := sec.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
