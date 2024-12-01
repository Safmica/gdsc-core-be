package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvInit() error{
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}
	return nil
}

func DSNInit() (string, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
)

    if dsn == "" {
        return "", errors.New("DB_PATH environment variable is not set")
    }
	return dsn, nil
}

func SecretKeyInit() (string, error){
	secretKey := os.Getenv("SECRET_KEY")
    if secretKey == "" {
        return "", errors.New("SECRET_KEY environment variable is not set")
    }
	return secretKey, nil
}

func TemplateInit() (string, error){
	template := os.Getenv("TEMPLATE_PATH")
	if template == "" {
		return "", errors.New("TEMPLATE_PATH environment variable is not set")
	}
	return template, nil
}