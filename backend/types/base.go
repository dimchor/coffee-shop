package types

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint64 `gorm:"primaryKey;not null;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
