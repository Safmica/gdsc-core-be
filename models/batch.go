package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Batch struct {
	IDBatch     uuid.UUID   `gorm:"type:char(36);primaryKey;column:id_batch" json:"id_batch"`
	Year        int         `gorm:"type:year;column:year" json:"year"`
	Members     []Member    `gorm:"foreignKey:IDBatch" json:"members"`
	Divisions   []Division  `gorm:"foreignKey:IDBatch" json:"divisions"`
	Activities  []Activity  `gorm:"foreignKey:IDBatch" json:"activities"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Batch) TableName() string {
	return "batch"
}
