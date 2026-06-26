package userdmn

import "time"

type User struct {
	Name       UserName
	CreateTime time.Time
	UpdateTime time.Time
	FirstName  string
	LastName   string
	Email      string
	State      UserState
}
