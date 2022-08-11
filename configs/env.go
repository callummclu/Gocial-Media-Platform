package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvByName(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv(name)
}

func EnvPORT() string {
	return ":" + getEnvByName("PORT")
}
