package model

import "time"

type Comment struct {
	Id int
	From string
	To string
	Time time.Time
	Content string
	Cid int
}
