package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUsers create user in database
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user model.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	// Validates if the fields are filled
	if err := user.Prepare("register"); err != nil {
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
	user.ID, err = repository.CreateUsers(user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, &model.User{
		ID:   user.ID,
		Name: user.Name,
		Nick: user.Nick,
		Mail: user.Mail,
	})

}

// SearchUsers return list of users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	// we use the "user" to perform the search in the URL. (Ex: /usuarios?user=name.user)
	nameORNick := strings.ToLower(r.URL.Query().Get("user"))
	// if nameORNick == "" {
	// 	response.Erro(w, http.StatusBadRequest, errors.New("search parameter not informed"))
	// 	return
	// }

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	users, err := repository.SearchUsers(nameORNick)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if len(users) == 0 {
		response.Erro(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	response.JSON(w, http.StatusOK, users)

}

// SearchUser return specific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
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
	user, err := repository.SearchUser(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		response.Erro(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

// UpdateUser update specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get path params
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("You don't have permission to update this user"))
		return
	}

	// Perform body reading
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

	if err = user.Prepare("edit"); err != nil {
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
	err = repository.UpdateUser(userID, user)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser remove specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		response.Erro(w, http.StatusForbidden, errors.New("You don't have permission to delete this user"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	if erro := repository.DeleteUser(userID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
