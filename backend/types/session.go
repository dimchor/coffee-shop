package types

type Session struct {
	Base
	UserID uint64
	User   User
}
