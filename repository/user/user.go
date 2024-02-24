package user

import "github.com/mercy34mercy/go-http-server/model/user"

type UserRepository interface {
	GetUserById(id user.UserID) (*user.User, error)
	Save(user user.User) error
}
