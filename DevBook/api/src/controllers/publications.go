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
	"strconv"

	"github.com/gorilla/mux"
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

	if err = publication.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

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
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublications(db)
	publications, err := repository.SearchPublications(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)

}

func SearchPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublications(db)
	publication, err := repository.SearchPublicationID(publicationID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
