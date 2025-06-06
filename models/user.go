package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID
	FullName       string
	Email          string `gorm:"unique"`
	Password       string `json:"-"`
	Country        string
	StudentProfile *Student `gorm:"foreignKey:ProfileID"`
	TeacherProfile *Teacher `gorm:"foreignKey:ProfileID"`
	Role           string
	ProfilePicture string `gorm:"default:'https://avatar.iran.liara.run/public/girl'"`
}
