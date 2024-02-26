package types

import (
	"github.com/google/uuid"
)

const SESSION_ID_SIZE = 16

type Session struct {
	Base
	Token  [SESSION_ID_SIZE]byte `gorm:"unique;size:16"`
	UserID uint64
	User   User
}

func NewSession(user_id uint64) (uuid.UUID, *Session) {
	token := uuid.New()
	return token, &Session{Base{}, token, user_id, User{}}
}
