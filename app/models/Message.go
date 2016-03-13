package models

import (
	"time"
)

type Message struct {
	ID            uint
	UserID        uint
	ChatSessionID uint
	Text          string
	FileUrl       string
	FileMimeType  string
	SendAt        time.Time
}
