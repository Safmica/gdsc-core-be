package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participant struct {
	IDParticipant uuid.UUID `gorm:"type:char(36);primaryKey;column:id_participant" json:"id_participant"`
	IDMember      uuid.UUID `gorm:"type:char(36);column:id_member" json:"id_member"`
	IDActivity    uuid.UUID `gorm:"type:char(36);column:id_activity" json:"id_activity"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Participant) TableName() string {
	return "participant"
}
