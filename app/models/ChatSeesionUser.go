package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ChatSessionUser struct {
	gorm.Model
	ChatSessionID uint
	UserID        uint
	JoinedAt      time.Time
}
