package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    uint   `gorm:"primaryKey" json:"user_id"`
	Username  string `json:"username" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(100)"`
	Address   string `json:"address" gorm:"type:varchar(100)"`
	Phone     int    `json:"phone"`
	Password  string `json:"password" gorm:"type:varchar(12)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
