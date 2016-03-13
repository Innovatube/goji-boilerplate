package models

type ChatSession struct {
	ID                uint
	Name              string
	Topic             string
	ChatSessionTypeID uint
	OwnerID           uint
}
