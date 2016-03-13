package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ChatSessionStatus struct {
	gorm.Model
	Status    string
	CreatedAt time.Time
}
