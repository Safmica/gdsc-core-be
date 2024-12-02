package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FinalProject struct {
	IDFinalProject uuid.UUID `gorm:"type:char(36);primaryKey;column:id_final_project" json:"id_final_project"`
	IDMember       uuid.UUID `gorm:"type:char(36);column:id_member;foreignkey:IDMember" json:"id_member"`
	Title          string    `gorm:"type:varchar(100);column:title" json:"title"`
	Description    string    `gorm:"type:text;column:description" json:"description"`
	Url            string    `gorm:"type:varchar(255);column:url" json:"url"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (FinalProject) TableName() string {
	return "final_project"
}
