package model

import "time"

type Content struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Time time.Time `sql:"default now()" json:"time"`
	Author int `json:"author"`
	Views int `sql:"default 0" json:"views"`
	Category string `json:"category"`
	IsKey bool `sql:"default false" json:"is_key"`
	Content string `json:"content"`
}

const (
	SHARE = "SHARE"
	QUESTION = "QUESTION"
	RECRUIMENT = "RECRUIMENT"
	TOPIC = "RECRUIMENT"
)