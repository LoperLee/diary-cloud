package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	*gorm.Model
	Id        int
	Name      string
	CreatTime time.Time
}

// get users in database
func GetUsers() []Users {
	return []Users{}
}
