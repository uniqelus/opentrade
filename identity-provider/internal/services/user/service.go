package usersrv

import (
	"context"
	"time"

	userdmn "github.com/uniqelus/opentrade/identity-provider/internal/domains/user"
	userapi "github.com/uniqelus/opentrade/identity-provider/internal/transport/grpc/user"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *userdmn.User) error
}

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

var _ userapi.UserService = &Service{}

// CreateUser implements [userapi.UserService].
func (s *Service) CreateUser(ctx context.Context, opts ...userdmn.CreateUserOption) (*userdmn.User, error) {
	options, err := userdmn.NewCreateUserOptions(opts...)
	if err != nil {
		return nil, err
	}

	userName, err := userdmn.NewName()
	if err != nil {
		return nil, err
	}

	user := &userdmn.User{
		Name:       userName,
		FirstName:  options.FirstName,
		LastName:   options.LastName,
		Email:      options.Email,
		State:      userdmn.UserStateActive,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := s.userRepository.SaveUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
