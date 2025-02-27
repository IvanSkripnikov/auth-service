package models

import (
	"os"
	"strconv"
)

type Config struct {
	Database          Database
	BillingServiceUrl string
}

func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: Database{
			Address:  os.Getenv("DB_ADDRESS"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DB:       os.Getenv("DB_NAME"),
		},
		BillingServiceUrl: os.Getenv("BILLING_SERViCE_URL"),
	}, nil
}

func GetRequiredVariables() []string {
	return []string{
		// Обязательные переменные окружения для подключения к БД сервиса
		"DB_ADDRESS", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME",

		// Обязательные переменные для редиректа в сервис платежей
		"BILLING_SERViCE_URL",
	}
}
