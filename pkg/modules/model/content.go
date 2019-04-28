package model

import "time"

type Content struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Time time.Time `gorm:"default:now()" json:"time"`
	Author int `json:"author"`
	Views int `gorm:"default:0" json:"views"`
	Category string `json:"category"`
	IsKey bool `gorm:"default:false" json:"is_key"`
	Content string `gorm:"text" json:"content"`
}

const (
	SHARE = "SHARE"
	QUESTION = "QUESTION"
	RECRUIMENT = "RECRUIMENT"
	TOPIC = "RECRUIMENT"
)