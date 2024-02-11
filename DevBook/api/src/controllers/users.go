package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/response"
	"api/src/sec"
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
		response.Erro(w, http.StatusForbidden, errors.New("you don't have permission to update this user"))
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
		response.Erro(w, http.StatusForbidden, errors.New("you don't have permission to delete this user"))
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

// FollowUser follow specific user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	followID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	// A user cannot follow himself
	if followID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Populate the followers table, passing the userID and followeID
	repository := repository.NewRepositoryUser(db)
	if err = repository.Follow(userID, followID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser unfollow specific user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	followID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		response.Erro(w, http.StatusForbidden, errors.New("it is not possible to be follow yourself"))
		return
	}

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	if err = repository.UnFollow(userID, followID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// SearchFollowers Get all followers of a specific user
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repository.NewRepositoryUser(db)
	followers, err := repository.SearchFollowers(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, followers)
}

// SearchFollowing Get all following of a specific user
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
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
	users, err := repository.SearchFollowing(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if len(users) < 1 {
		response.JSON(w, http.StatusNoContent, nil)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

// UpdatePassword change a specific user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := authentication.ExtractUserID(r)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if userIDToken != userID {
		response.Erro(w, http.StatusForbidden, errors.New("it is not allowed to update the password of a user other than your own"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var password model.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
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
	passwordSavedDatabase, err := repository.SearchPassword(userID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = sec.CheckPassword(passwordSavedDatabase, password.Current); err != nil {
		response.Erro(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	passwordHash, err := sec.Hash(password.New)
	if err != nil {
		response.Erro(w, http.StatusBadGateway, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(passwordHash)); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
