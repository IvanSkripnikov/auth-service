package main

import (
	"fmt"

	"authenticator/helpers"
	"authenticator/httphandler"
	"authenticator/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Config error: %v", err))
	}
	helpers.InitConfig(config)

	// настройка коннекта к БД
	_, err = helpers.InitDataBase(config.Database)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cant initialize DB: %v", err))
	}

	// инициализация сессий
	helpers.SessionsMap = map[string]models.User{}

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
