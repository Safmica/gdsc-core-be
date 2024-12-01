package models

import (
	"time"

	"gorm.io/gorm"
)

type Configure struct {
	ID           int            `gorm:"type:int;autoIncrement;column:id" json:"id"`
	CurrentBatch int            `gorm:"type:year;column:current_batch" json:"current_batch"`
	CreatedAt    time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Configure) TableName() string {
	return "configure"
}
