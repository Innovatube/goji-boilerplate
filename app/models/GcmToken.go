package models

import (
	"github.com/jinzhu/gorm"
)

type GcmToken struct {
	gorm.Model
	Token string
}
