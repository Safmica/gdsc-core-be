package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Division struct {
	IDDivision    uuid.UUID `gorm:"type:char(36);primaryKey;column:id_division" json:"id_division"`
	IDBatch       uuid.UUID `gorm:"type:char(36);column:id_batch" json:"id_batch"`
	Name          string    `gorm:"type:varchar(100);column:name" json:"name"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Division) TableName() string {
	return "division"
}
