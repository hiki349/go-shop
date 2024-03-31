package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GqlPort         string
	RestPort        string
	ConnStrPostgres string
	ConnStrMongo    string
	JwtSecret       string
	Mode            string
}

func MustGetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	gqlPort := os.Getenv("GQL_PORT")
	restPort := os.Getenv("REST_PORT")
	connStrPostgres := os.Getenv("CONNECTION_STRING_POSTGRES")
	connStrMongo := os.Getenv("CONNECTION_STRING_MONGO")
	jwtSecret := os.Getenv("JWT_SECRET")
	mode := os.Getenv("MODE")

	return Config{
		GqlPort:         gqlPort,
		RestPort:        restPort,
		ConnStrPostgres: connStrPostgres,
		ConnStrMongo:    connStrMongo,
		JwtSecret:       jwtSecret,
		Mode:            mode,
	}
}
