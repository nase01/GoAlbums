package models

type Album struct {
	CtmBasemModel
	Title  string  `json:"title" gorm:"column:title;unique;not null;varchar(255);index"`
	Artist string  `json:"artist" gorm:"column:artist;not null;varchar(255);"`
	Prices float32 `json:"price" gorm:"column:price;not null;double;"`
}
