package usecase

import (
	"fmt"

	user_model "github.com/mercy34mercy/go-http-server/model/user"
	"github.com/mercy34mercy/go-http-server/repository/user"
	"github.com/mercy34mercy/go-http-server/repository/user/psql"
)

type GetUserByID interface {
	Execute(id user_model.UserID) (*user_model.User, error)
}

var _ GetUserByID = (*getUserByIDImpl)(nil)

type getUserByIDImpl struct {
	userRepository user.UserRepository
}

func (u *getUserByIDImpl) Execute(id user_model.UserID) (*user_model.User, error) {
	user, err := u.userRepository.GetUserById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func NewGetUserByID(repo psql.UserRepository) GetUserByID {
	return &getUserByIDImpl{
		userRepository: &repo,
	}
}
