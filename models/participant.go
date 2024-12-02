package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participant struct {
	IDParticipant uuid.UUID `gorm:"type:char(36);primaryKey;column:id_participant" json:"id_participant"`
	IDMember      uuid.UUID `gorm:"type:char(36);column:id_member" json:"id_member"`
	IDActivity    uuid.UUID `gorm:"type:char(36);column:id_activity" json:"id_activity"`
	gorm.Model
}

func (Participant) TableName() string {
	return "participant"
}
