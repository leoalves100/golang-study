package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/sec"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	userAndMailSaved, err := repository.SearchUserMail(user.Mail)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	if err := sec.CheckPassword(userAndMailSaved.Password, user.Password); err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("Logou porra"))

}
