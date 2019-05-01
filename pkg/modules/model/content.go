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
	Content string `gorm:"type:text" json:"content"`
}

const (
	SHARE = "Share"
	QUESTION = "Q&A"
	RECRUIMENT = "Recruit"
	TOPIC = "Topic"
)

type Category struct {
	Id int `json:""`
	Name string `json:"name"`
	Time time.Time `gorm:"default:now()" json:"time"`
	Creator int `json:"creator"`
	Alias string `json:"alias"`
	Weight int `json:"weight"`
}