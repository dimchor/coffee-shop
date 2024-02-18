package types

import "crypto/rand"

const SESSION_ID_SIZE = 16

type Session struct {
	Base
	Token  [SESSION_ID_SIZE]byte `gorm:"unique;size:16"`
	UserID uint64
	User   User
}

func NewSession(user_id uint64) (*Session, error) {
	session_id := make([]byte, SESSION_ID_SIZE)

	if _, err := rand.Read(session_id); err != nil {
		return nil, err
	}

	return &Session{Base{}, [16]byte(session_id), user_id, User{}}, nil
}
