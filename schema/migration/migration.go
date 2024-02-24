package main

import (
	"database/sql"
	"fmt"

	"github.com/mercy34mercy/go-http-server/log"

	_ "github.com/lib/pq"

	"github.com/mercy34mercy/go-http-server/config"
)

// user.sqlをmigration

func main() {
	client, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.POSTGRESQL_HOST(), config.POSTGRESQL_PORT(), config.POSTGRESQL_USER_NAME(), config.POSTGRESQL_PASSWORD(), config.POSTGRESQL_DB_NAME()))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database: %w", err))
	}
	defer client.Close()

	_, err = client.Exec(`
		DROP TABLE IF EXISTS users;
	`)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to drop table: %w", err))
	}

	// user.sqlをmigration
	_, err = client.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(36) PRIMARY KEY,
			age INT,
			name VARCHAR(255)
		);
	`)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

}
