package models

type User struct {
	CtmBasemModel
	Email    string `json:"email" gorm:"column:email;unique;index"`
	Password string `json:"-" gorm:"column:password;varchar(255);"` // Exclude Password from JSON Output
	FullName string `json:"fullname" gorm:"column:fullname;not null;varchar(255);"`
}
