package main

import (
	"database/sql"
	"fmt"
	"time"

	"net/http"

	_ "github.com/lib/pq"
	"github.com/mercy34mercy/go-http-server/config"
	"github.com/mercy34mercy/go-http-server/handler/health"
	v1 "github.com/mercy34mercy/go-http-server/handler/v1"
	"github.com/mercy34mercy/go-http-server/log"
	"github.com/mercy34mercy/go-http-server/repository/user/psql"
	"github.com/mercy34mercy/go-http-server/usecase"
)

func main() {
	// psqlのcliを生成
	client, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.POSTGRESQL_HOST(), config.POSTGRESQL_PORT(), config.POSTGRESQL_USER_NAME(), config.POSTGRESQL_PASSWORD(), config.POSTGRESQL_DB_NAME()))
	if err != nil {
		log.Fatal(err, fmt.Errorf("failed to connect to database: %w", err))
	}
	defer client.Close()

	// 接続プールの設定
	client.SetMaxOpenConns(400)
	client.SetMaxIdleConns(400)
	client.SetConnMaxLifetime(60 * time.Minute)

	userRepo := psql.NewUserRepository(client)

	http.HandleFunc("GET /", health.HealthHandler())
	http.HandleFunc("GET /health", health.HealthHandler())
	http.HandleFunc("POST /user", v1.CreateUserHandler(usecase.NewSaveUser(*userRepo)))
	http.HandleFunc("GET /user/{user_id}", v1.GetUserHandler(usecase.NewGetUserByID(*userRepo)))
	http.ListenAndServe(":8080", nil)
}
