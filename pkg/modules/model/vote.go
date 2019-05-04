package model

type Vote struct {
	Id int `json:"id"`
	Creator int `json:"creator"`
	JoinNum int `json:"participant"`
	Title string `json:"title"`

}
