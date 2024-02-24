package usecase

import (
	"fmt"

	user_model "github.com/mercy34mercy/go-http-server/model/user"
	"github.com/mercy34mercy/go-http-server/repository/user"
	"github.com/mercy34mercy/go-http-server/repository/user/psql"
)

type SaveUser interface {
	Execute(user user_model.User) error
}

var _ SaveUser = (*saveUserImpl)(nil)

type saveUserImpl struct {
	userRepository user.UserRepository
}

func (s *saveUserImpl) Execute(user user_model.User) error {
	err := s.userRepository.Save(user)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func NewSaveUser(repo psql.UserRepository) SaveUser {
	return &saveUserImpl{
		userRepository: &repo,
	}
}
