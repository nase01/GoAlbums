package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
	RoleSuper Role = "super"
)

type User struct {
	CtmBasemModel
	Email       string `json:"email" gorm:"column:email;unique;index"`
	Password    string `json:"-" gorm:"column:password;varchar(255);"` // Exclude Password from JSON Output
	FullName    string `json:"fullname" gorm:"column:fullname;not null;varchar(255);"`
	Role        Role   `json:"role" gorm:"column:role;type:enum('user', 'admin', 'super');default:'user'"`
	IPWhitelist string `json:"ipWhitelist" gorm:"column:ipWhitelist;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Role == "" {
		user.Role = RoleUser
	}
	return
}
