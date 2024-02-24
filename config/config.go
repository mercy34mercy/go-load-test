package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func POSTGRESQL_USER_NAME() string {
	v, ok := os.LookupEnv("POSTGRESQL_USER_NAME")
	if !ok {
		log.Fatal("POSTGRES_USER_NAME is not set")
	}
	return v
}

func POSTGRESQL_PASSWORD() string {
	v, ok := os.LookupEnv("POSTGRESQL_PASSWORD")
	if !ok {
		log.Fatal("POSTGRES_PASSWORD is not set")
	}
	return v
}

func POSTGRESQL_DB_NAME() string {
	v, ok := os.LookupEnv("POSTGRESQL_DB_NAME")
	if !ok {
		log.Fatal("POSTGRES_DB_NAME is not set")
	}
	return v
}

func POSTGRESQL_HOST() string {
	v, ok := os.LookupEnv("POSTGRESQL_HOST")
	if !ok {
		log.Fatal("POSTGRES_HOST is not set")
	}
	return v
}

func POSTGRESQL_PORT() string {
	v, ok := os.LookupEnv("POSTGRESQL_PORT")
	if !ok {
		log.Fatal("POSTGRES_PORT is not set")
	}
	return v
}

func POSTGRESQL_MAX_POOL_CONNS() string {
	v, ok := os.LookupEnv("POSTGRESQL_MAX_POOL_CONNS")
	if !ok {
		log.Fatal("POSTGRES_MAX_POOL_CONNS is not set")
	}
	return v
}
