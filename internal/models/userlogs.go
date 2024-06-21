package models

type UserLogs struct {
	CtmBasemModel
	UserId   string `json:"userId" gorm:"column:userId;unique;not null;varchar(255);index"`
	Activity string `json:"activity" gorm:"column:activity;type:text;"`
}
