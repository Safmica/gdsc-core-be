package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	IDMember     uuid.UUID    `gorm:"type:char(36);primaryKey;column:id_member" json:"id_member"`
	IDUser       uuid.UUID    `gorm:"type:char(36);column:id_user" json:"id_user"`
	IDRole       uuid.UUID    `gorm:"type:char(36);column:id_role" json:"id_role"`
	IDBatch      uuid.UUID    `gorm:"type:char(36);column:id_batch" json:"id_batch"`
	Email        string       `gorm:"-" json:"email"`
	IDDivision   uuid.UUID    `gorm:"type:char(36);column:id_division" json:"id_division"`
	Activities   []Activity   `gorm:"many2many:participant;foreignKey:IDMember;joinForeignKey:IDMember;References:IDActivity;joinReferences:IDActivity" json:"activities"`
	FinalProject FinalProject `gorm:"foreignKey:IDMember" json:"final_project"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Member) TableName() string {
	return "member"
}
