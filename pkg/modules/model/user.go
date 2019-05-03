package model

import "time"

type User struct {
	Id int `gorm:"primary_key" json:"id"`
	Time time.Time `gorm:"default:now()" json:"time"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	Avatar string `json:"avatar"`
	Grades int `json:"grades"`
	Description string `json:"description"`
	Mail string `json:"mail"`
	Sex int `json:"sex"`
	Favorites []Favorite `json:"favorites"`
	Role string `gorm:"default:'user'" json:"role"`
}

type Favorite struct {
	Id int `gorm:"primary_key" json:"id"`
	Uid int `json:"uid"`
	Cid int `json:"cid"`
	Valid bool `json:"valid"`
}
