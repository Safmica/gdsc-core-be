package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	IDUser         uuid.UUID `gorm:"type:char(36);primaryKey;column:id_user" json:"id_user"`
	Password       string    `gorm:"type:varchar(255);column:password" json:"password"`
	NewPassword    string    `gorm:"-" json:"new_password"`
	Name           string    `gorm:"type:varchar(100);column:name" json:"name"`
	Email          string    `gorm:"type:varchar(100);unique;column:email" json:"email"`
	University     string    `gorm:"type:varchar(100);column:university" json:"university"`
	Major          string    `gorm:"type:varchar(50);column:major" json:"major"`
	Year           int       `gorm:"type:year;column:year" json:"year"`
	NIM            string    `gorm:"type:varchar(20);column:nim" json:"nim"`
	CurrentBatch   int       `gorm:"type:year;column:current_batch" json:"current_batch"`
	BatchHistories []Batch   `gorm:"many2many:user_batches;" json:"batch_histories"`

	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
