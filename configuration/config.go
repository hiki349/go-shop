package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	ConnStr string
}

func MustGetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	connStr := os.Getenv("CONNECTION_STRING")

	return Config{
		Port:    port,
		ConnStr: connStr,
	}
}
