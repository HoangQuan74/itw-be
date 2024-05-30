package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailConversationEnity struct {
	gorm.Model
	SenderID       uuid.UUID `gorm:"size:255"`
	Content        string    `gorm:"type:longtext"`
	ConversationID uuid.UUID `gorm:"size:255"`

	Sender       UserPerson
	Conversation Conversation
}
