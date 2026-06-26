package userdmn

type UserState int

const (
	UserStateUnspecified UserState = iota
	UserStateActive
	UserStateSuspended
	UserStateLocked
)
