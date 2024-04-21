package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailConversationEnity struct {
	gorm.Model
	SenderID       uuid.UUID
	Content        string `gorm:"type:longtext"`
	ConversationID uuid.UUID

	Sender       UserPerson
	Conversation Conversation
}
