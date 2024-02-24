package psql_test

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/go-cmp/cmp"
	user_model "github.com/mercy34mercy/go-http-server/model/user"
	"github.com/mercy34mercy/go-http-server/repository/user/psql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbName     = "test"
	dbUser     = "test"
	dbPassword = "test"
)

func TestSaveAndFetch(t *testing.T) {
	cases := []struct {
		name string
		user user_model.User
		want user_model.User
	}{
		{
			name: "success",
			user: user_model.User{
				ID:   "ab12",
				Name: "test",
				Age:  20,
			},
			want: user_model.User{
				ID:   "ab12",
				Name: "test",
				Age:  20,
			},
		},
	}
	ctx := context.Background()
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("../../../schema/", "database", "user.sql")),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	repo := psql.NewUserRepository(client)

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := repo.Save(c.user)
			if err != nil {
				t.Fatalf("failed to save user: %s", err)
			}

			got, err := repo.GetUserById(c.user.ID)
			if err != nil {
				t.Fatalf("failed to get user by id: %s", err)
			}

			if diff := cmp.Diff(*got, c.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
