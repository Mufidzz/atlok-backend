package structs

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	VerifiedAt *time.Time
	Username   string `gorm:"unique"`
	Password   string
	Access     uint
}
