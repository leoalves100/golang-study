package controllers

import (
	"api/src/authentication"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"crud/banco"
	"encoding/json"
	"io"
	"net/http"
)

// CreatePublication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication model.Publication
	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	db, err := banco.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublications(db)

	publication.ID, err = repository.CreatePublication(publication)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, publication)
}

func SearchPublications(w http.ResponseWriter, r *http.Request) {

}

func SearchPublication(w http.ResponseWriter, r *http.Request) {

}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
