package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"authenticator/models"

	"github.com/IvanSkripnikov/go-logger"
)

var SessionsMap map[string]models.User

func Register(w http.ResponseWriter, r *http.Request) {
	category := "/register"
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if checkError(w, err, category) {
		return
	}

	user.Created = GetCurrentDate()
	user.Updated = GetCurrentDate()
	user.Active = 1
	user.CategoryID = 1

	// create user
	var data ResponseData
	var httpStatus int

	err = GormDB.Create(&user).Error
	if checkError(w, err, category) {
		return
	}

	// create user account
	err = createUserAccount(user.ID)
	if err != nil {
		logger.Errorf("Can't register user account: %v", err)
		data = ResponseData{
			"response": err.Error(),
		}
		httpStatus = http.StatusBadRequest
	} else {
		data = ResponseData{
			"response": user.ID,
		}
		httpStatus = http.StatusOK
	}

	SendResponse(w, data, category, httpStatus)
}

func Login(w http.ResponseWriter, r *http.Request) {
	category := "/login"
	// handle incoming params
	var rp models.LoginParams
	err := json.NewDecoder(r.Body).Decode(&rp)
	if checkError(w, err, category) {
		return
	}

	var data ResponseData
	var httpStatus int
	var user models.User

	err = GormDB.Where(&models.User{Username: rp.Username, Password: rp.Password}).First(&user).Error
	if checkError(w, err, category) {
		return
	}
	user.Password = ""

	// записываем сессию
	sessionID := GenerateSessionID(user)
	SessionsMap[sessionID] = user

	// задаём cookie
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	// возвращаем ответ
	data = ResponseData{
		"response": models.Success,
	}
	httpStatus = http.StatusOK

	SendResponse(w, data, category, httpStatus)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	var data ResponseData
	var httpStatus int

	cookie, err := r.Cookie("session_id")
	if err != nil {
		data = ResponseData{
			"response": "Error: " + err.Error(),
		}
		httpStatus = http.StatusBadRequest
	}

	if cookie.Value != "" {
		val, ok := SessionsMap[cookie.Value]
		if ok {
			data = ResponseData{
				"data": val,
			}
			httpStatus = http.StatusOK
			w.Header().Add("X-UserId", strconv.Itoa(val.ID))
			w.Header().Add("X-User", val.Username)
			w.Header().Add("X-Email", val.Email)
			w.Header().Add("X-First-Name", val.FirstName)
			w.Header().Add("X-Last-Name", val.LastName)
		} else {
			data = ResponseData{
				"error": "Not authorized",
			}
			httpStatus = http.StatusUnauthorized
		}
	} else {
		data = ResponseData{
			"error": "Not authorized",
		}
		httpStatus = http.StatusUnauthorized
	}

	SendResponse(w, data, "/auth", httpStatus)
}

func SignIn(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"message": "Please go to login and provide Login/Password",
	}
	SendResponse(w, data, "/signin", http.StatusOK)
}

func Logout(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"status": "OK",
	}
	// задаём cookie
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now(),
	}
	http.SetCookie(w, cookie)

	SendResponse(w, data, "/logout", http.StatusOK)
}

func Sessions(w http.ResponseWriter, _ *http.Request) {
	data := ResponseData{
		"sessions": SessionsMap,
	}
	SendResponse(w, data, "/sessions", http.StatusOK)
}

func createUserAccount(id int) error {
	newAccount := models.Account{UserID: id, Balance: 0}
	jsonData, err := json.Marshal(newAccount)
	if err != nil {
		return err
	}
	// Отправляем POST-запрос
	resp, err := http.Post(Config.BillingServiceUrl+"/v1/account/create", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
