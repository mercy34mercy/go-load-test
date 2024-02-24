package psql

import (
	"database/sql"
	"fmt"

	"github.com/mercy34mercy/go-http-server/model/user"
)

type UserRepository struct {
	Cli *sql.DB
}

func (r *UserRepository) GetUserById(id user.UserID) (*user.User, error) {
	row := r.Cli.QueryRow("SELECT * FROM users WHERE id = $1", id.String())
	var u user.User
	err := row.Scan(&u.ID, &u.Age, &u.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}
	return &u, nil
}

func (r *UserRepository) Save(user user.User) error {
	_, err := r.Cli.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", user.ID.String(), user.Name, user.Age)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func NewUserRepository(cli *sql.DB) *UserRepository {
	return &UserRepository{Cli: cli}
}
