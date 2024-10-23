package env

import (
	"errors"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type EnvReader struct {
	DB_URL     string
	JWT_SECRET []byte
}

var (
	once   sync.Once
	reader *EnvReader
	err    error
)

func loadENV() (*EnvReader, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("can't load .env file")
	}

	dbURL := os.Getenv("DB_URI")
	jwtSecret := os.Getenv("JWT_SECRET")

	if dbURL == "" {
		return nil, errors.New("DB_URI is not present in ENV variables")
	}

	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET is not present in ENV variables")
	}

	return &EnvReader{
		DB_URL:     dbURL,
		JWT_SECRET: []byte(jwtSecret),
	}, nil

}

func GetENVReader() (*EnvReader, error) {
	once.Do(func() {
		reader, err = loadENV()
	})
	return reader, err
}
