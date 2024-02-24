package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/mercy34mercy/go-http-server/log"

	_ "github.com/lib/pq"

	"github.com/mercy34mercy/go-http-server/config"
)

func main() {
	// データベース接続設定
	client, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.POSTGRESQL_HOST(), config.POSTGRESQL_PORT(), config.POSTGRESQL_USER_NAME(), config.POSTGRESQL_PASSWORD(), config.POSTGRESQL_DB_NAME()))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database: %w", err))
	}
	defer client.Close()

	// テーブルの再作成
	_, err = client.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to drop table: %w", err))
	}

	_, err = client.Exec(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		age INT,
		name VARCHAR(255)
	);`)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create table: %w", err))
	}

	// 1000人のユーザーを追加
	for i := 0; i < 1000; i++ {
		id := uuid.New().String()          // UUIDを生成
		age := rand.Intn(100) + 1          // 1 から 100 までの年齢をランダムに生成
		name := fmt.Sprintf("User%d", i+1) // User1, User2, ..., User1000 の名前を生成

		_, err = client.Exec(`INSERT INTO users (id, age, name) VALUES ($1, $2, $3);`, id, age, name)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to insert user: %w", err))
		}
	}
}
