package service

import (
	"errors"
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
)

type ContentService struct {
	client *mysql.Client
}

func NewContentService(client *mysql.Client) *ContentService {
	return &ContentService{
		client: client,
	}
}

func (c *ContentService) ValidateContent(content model.Content) error {
	if content.Title == "" {
		return errors.New("title empty")
	} else if content.Content == "" {
		return errors.New("content empty")
	}
	return nil
}

func (c *ContentService) GetContentById(id int) (content model.Content, err error) {
	err = c.client.DB.Table("contents").Where("id = ?", id).Scan(&content).Error
	return
}

func (c *ContentService) CreateContent(content model.Content) error {
	if content.Title == "" || content.Content == "" {
		return errors.New("title and content show not be empty")
	}
	return c.client.DB.Table("contents").Create(&content).Error
}