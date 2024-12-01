package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	IDUser       uuid.UUID `gorm:"type:char(36);column:id_user;primaryKey" json:"id_user"`
	RefreshToken string    `gorm:"type:varchar(255);column:refresh_token" json:"refresh_token"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Token) TableName() string {
	return "token"
}