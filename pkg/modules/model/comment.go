package model

import "time"

type Comment struct {
	Id int `gorm:"primary_key" json:"id"`
	From string `gorm:"varchar" json:"from"`
	To string `gorm:"varchar" json:"to"`
	Time time.Time `gorm:"default:now()" json:"time"`
	Content string `gorm:"text" json:"content"`
	Cid int `json:"cid"`
}
