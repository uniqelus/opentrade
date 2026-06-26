package userrepo

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	userdmn "github.com/uniqelus/opentrade/identity-provider/internal/domains/user"
)

type dbUserState string

const (
	dbStateActive    dbUserState = "USER_STATE_ACTIVE"
	dbStateSuspended dbUserState = "USER_STATE_SUSPENDED"
	dbStateLocked    dbUserState = "USER_STATE_LOCKED"
)

type userDTO struct {
	ID         uuid.UUID
	Email      string
	FirstName  string
	LastName   string
	State      dbUserState
	CreateTime time.Time
	UpdateTime time.Time
}

func toDTO(u *userdmn.User) (userDTO, error) {
	uid := u.Name.ID()
	if uid == uuid.Nil {
		return userDTO{}, fmt.Errorf("invalid user resource name or empty uuid")
	}

	return userDTO{
		ID:         uid,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		State:      mapDomainStateToDB(u.State),
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
	}, nil
}

func mapDomainStateToDB(s userdmn.UserState) dbUserState {
	switch s {
	case userdmn.UserStateActive:
		return dbStateActive
	case userdmn.UserStateSuspended:
		return dbStateSuspended
	case userdmn.UserStateLocked:
		return dbStateLocked
	default:
		return dbStateActive
	}
}
