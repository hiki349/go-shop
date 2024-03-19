package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GqlPort   string
	RestPort  string
	ConnStr   string
	JwtSecret string
}

func MustGetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	gqlPort := os.Getenv("GQL_PORT")
	restPort := os.Getenv("REST_PORT")
	connStr := os.Getenv("CONNECTION_STRING")
	jwtSecret := os.Getenv("JWT_SECRET")

	return Config{
		GqlPort:   gqlPort,
		RestPort:  restPort,
		ConnStr:   connStr,
		JwtSecret: jwtSecret,
	}
}
