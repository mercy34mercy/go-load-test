package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mercy34mercy/go-http-server/config"
	"github.com/mercy34mercy/go-http-server/handler/health"
	v1 "github.com/mercy34mercy/go-http-server/handler/v1"
	"github.com/mercy34mercy/go-http-server/repository/user/psql"
	"github.com/mercy34mercy/go-http-server/usecase"
)

func main() {
	// psqlのcliを生成
	client, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=%s", config.POSTGRESQL_HOST(), config.POSTGRESQL_PORT(), config.POSTGRESQL_USER_NAME(), config.POSTGRESQL_PASSWORD(), config.POSTGRESQL_DB_NAME(), config.POSTGRESQL_MAX_POOL_CONNS()))
	if err != nil {	
		log.Fatal("failed to connect to database %w", err)
	}
	defer client.Close()

	userRepo := psql.NewUserRepository(client)

	http.HandleFunc("GET /", health.HealthHandler())
	http.HandleFunc("GET /health", health.HealthHandler())
	http.HandleFunc("POST /user", v1.CreateUserHandler(usecase.NewSaveUser(*userRepo)))
	http.HandleFunc("GET /user/{user_id}", v1.GetUserHandler(usecase.NewGetUserByID(*userRepo)))
	http.ListenAndServe(":8080", nil)
}