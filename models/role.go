package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	IDRole    uuid.UUID      `gorm:"type:char(36);primaryKey;column:id_role" json:"id_role"`
	KodeRole  string         `gorm:"type:varchar(2);column:kode_role" json:"kode_role"`
	Nama      string         `gorm:"type:varchar(50);column:nama" json:"nama"`
	Members   []Member       `gorm:"foreignKey:IDRole" json:"members"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Role) TableName() string {
	return "role"
}

