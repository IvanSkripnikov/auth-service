package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"authenticator/models"
)

var Config *models.Config

func InitConfig(cfg *models.Config) {
	Config = cfg
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func FormatResponse(w http.ResponseWriter, httpStatus int, category string) {
	w.WriteHeader(httpStatus)

	data := ResponseData{
		"error": "Unsuccessfull request",
	}
	SendResponse(w, data, category, httpStatus)
}

func GenerateSessionID(user models.User) string {
	str := time.Now().Format("YYYY-mm-dd_H:i:s") + user.Username

	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
