package configuration

import (
	"fmt"
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

	dbUsername := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", dbUsername, dbPass, dbPort, dbName)

	return Config{
		Port:    port,
		ConnStr: connStr,
	}
}
