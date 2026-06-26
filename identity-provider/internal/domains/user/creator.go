package userdmn

import (
	"context"
)

type Creator interface {
	CreateUser(ctx context.Context, opts ...CreateUserOption) (*User, error)
}

type CreateUserOptions struct {
	FirstName string
	LastName  string
	Email     string
}

func NewCreateUserOptions(opts ...CreateUserOption) (*CreateUserOptions, error) {
	options := &CreateUserOptions{}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	return options, nil
}

type CreateUserOption func(*CreateUserOptions) error

func WithCreateUserFirstName(value string) CreateUserOption {
	return func(cuo *CreateUserOptions) error {
		cuo.FirstName = value
		return nil
	}
}

func WithCreateUserLastName(value string) CreateUserOption {
	return func(cuo *CreateUserOptions) error {
		cuo.LastName = value
		return nil
	}
}

func WithCreateUserEmail(value string) CreateUserOption {
	return func(cuo *CreateUserOptions) error {
		cuo.Email = value
		return nil
	}
}
