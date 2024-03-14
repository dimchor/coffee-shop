package types

import (
	"github.com/google/uuid"
)

const SESSION_ID_SIZE = 16

type Session struct {
	Base
	Token  string `gorm:"unique"`
	UserID uint64
	User   User
}

func NewSession(user_id uint64) (string, *Session) {
	token := uuid.NewString()
	return token, &Session{Base{}, token, user_id, User{}}
}
