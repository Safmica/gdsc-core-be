package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	IDActivity  uuid.UUID `gorm:"type:char(36);primaryKey;column:id_activity" json:"id_activity"`
	IDBatch     uuid.UUID `gorm:"type:char(36);column:id_batch" json:"id_batch"`
	Name        string    `gorm:"type:varchar(100);column:name" json:"name"`
	Description string    `gorm:"type:text;column:description" json:"description"`
	Date        time.Time `gorm:"type:timestamp;column:date" json:"date"`
	Members     []Member  `gorm:"many2many:participant;foreignKey:IDActivity;joinForeignKey:IDActivity;References:IDMember;joinReferences:IDMember" json:"members"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Activity) TableName() string {
	return "activity"
}
