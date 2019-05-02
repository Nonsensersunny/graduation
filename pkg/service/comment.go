package service

import (
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
)

type CommentService struct {
	client *mysql.Client
}

func NewCommentService(client *mysql.Client) *CommentService {
	return &CommentService{
		client: client,
	}
}

func (c *CommentService) CreateComment(comment model.Comment) error {
	return c.client.DB.Table("comments").Create(&comment).Error
}


func (c *CommentService) GetCommentsByCid(id int) (comments []model.Comment, err error) {
	err = c.client.DB.Table("comments").Where("cid = ?", id).Scan(&comments).Error
	return
}
