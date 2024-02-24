package v1_test

import (
	"testing"

	"github.com/mercy34mercy/go-http-server/model/user"
)

// wip
func TestCreateUser(t *testing.T) {
	cases := []struct {
		names          string
		user           user.User
		wantErr        error
		wantStatusCode int
	}{
		{
			names: "success",
			user: user.User{
				ID:   "1",
				Name: "test",
				Age:  20,
			},
			wantErr:        nil,
			wantStatusCode: 200,
		},
	}

	for _, c := range cases {
		t.Run(c.names, func(t *testing.T) {
			
		})
	}
}
