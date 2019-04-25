package model

import "time"

type User struct {
	Id int
	Time time.Time
	Username string
	Password string
	Salt string
	Avatar string
	Grades int
	Description string
	Mail string
	Sex int
	Favorites []Favorite
}

type Favorite struct {
	Id int
	Uid int
	Cid int
}
