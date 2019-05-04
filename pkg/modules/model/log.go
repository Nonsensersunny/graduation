package model

import "time"

type Operation int

const (
	ADD = 100
	DEL = 101
	UPD = 102
)

type Archive struct {
	Time time.Time `json:"time"`
	UserId int `json:"user_id"`
	Operation Operation `json:"operation"`
	Data string `json:"data"`
}
