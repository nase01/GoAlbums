package models

type User struct {
	CtmBasemModel
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	FullName string `json:"fullname" gorm:"column:fullname;not null;varchar(255);"`
}
