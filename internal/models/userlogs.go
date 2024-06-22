package models

type UserLogs struct {
	CtmBasemModel
	UserId   string `json:"userId" gorm:"column:userId;unique;not null;varchar(50);"`
	Activity string `json:"activity" gorm:"column:activity;type:text;"`
}
