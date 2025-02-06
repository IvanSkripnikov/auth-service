package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"authenticator/logger"
	"authenticator/models"
)

func GetUsersList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/users/list"
	var users []models.User

	query := "SELECT id, username, first_name, last_name, email, phone, created, updated, active FROM users WHERE active = 1"
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Created, &user.Updated, &user.Active); err != nil {
			logger.Error(err.Error())
			continue
		}
		users = append(users, user)
	}

	data := ResponseData{
		"data": users,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/get"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if user.ID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "SELECT id, username, first_name, last_name, email, phone, created, updated, active FROM users WHERE id = ? AND active = 1"
	rows, err := DB.Prepare(query)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(user.ID).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Created, &user.Updated, &user.Active)
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"data": user,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func AddLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"addLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/add-loyalty", http.StatusOK)
}

func RemoveLoyalty(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"removeLoyalty": "OK",
	}
	SendResponse(w, data, "/v1/users/remove-loyalty", http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/create"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	query := "INSERT INTO users (username, first_name, last_name, email, phone, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)"
	currentTimestamp := GetCurrentTimestamp()
	rows, err := DB.Query(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, currentTimestamp, currentTimestamp)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully created!",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/update"
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	currentTimestamp := GetCurrentTimestamp()
	query := "UPDATE users SET username = ?, first_name = ?, last_name = ?, email = ?, phone = ?, updated = ? WHERE id = ?"
	rows, err := DB.Query(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Phone, currentTimestamp, user.ID)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully updated!",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	category := "/v1/users/block"
	var user models.User

	user.ID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if user.ID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM users WHERE id = ?", user.ID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "DELETE FROM users WHERE id = ?"
	rows, err := DB.Query(query, user.ID)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"message": "User successfully removed!",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func ResetPassword(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"resetPassword": "OK",
	}
	SendResponse(w, data, "/v1/users/reset-password", http.StatusOK)
}

func GetStatistics(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"getStatistics": "OK",
	}
	SendResponse(w, data, "/v1/users/statistics", http.StatusOK)
}
