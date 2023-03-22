package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConectionDB = ""
	ExecutionPortAPI  = 0
)

// LoadingConfigs initialization vars configs
func LoadingConfigs() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ExecutionPortAPI, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		ExecutionPortAPI = 9000
	}

	StringConectionDB = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
