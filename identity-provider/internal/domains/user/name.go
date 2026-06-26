package userdmn

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const segment = "users/"

var ErrInvalidName = errors.New("invalid user resource name format")

type UserName string

func NewName() (UserName, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate uuid: %w", err)
	}
	return UserName(segment + id.String()), nil
}

func ParseName(s string) (UserName, error) {
	if !strings.HasPrefix(s, segment) {
		return "", ErrInvalidName
	}

	idStr := strings.TrimPrefix(s, segment)
	if _, err := uuid.Parse(idStr); err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidName, err)
	}

	return UserName(s), nil
}

func (n UserName) String() string {
	return string(n)
}

func (n UserName) ID() uuid.UUID {
	idStr := strings.TrimPrefix(string(n), segment)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func (n UserName) IsEmpty() bool {
	return n == ""
}
