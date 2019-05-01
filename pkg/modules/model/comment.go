package model

import "time"

type Comment struct {
	Id int `gorm:"primary_key" json:"id"`
	From int `gorm:"type:int(11)" json:"from"`
	To int `gorm:"type:int(11)" json:"to"`
	Time time.Time `gorm:"default:now()" json:"time"`
	Content string `gorm:"type:text" json:"content"`
	Cid int `json:"cid"`
}
