package service

import (
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
)

type LinkService struct {
	client *mysql.Client
}

func NewLinkService(client *mysql.Client) *LinkService {
	return &LinkService{
		client: client,
	}
}

func (l *LinkService) CreateLink(link model.Link) (err error) {
	err = l.client.DB.Table("links").Create(&link).Error
	return
}

func (l *LinkService) DeleteLink(id int) (err error) {
	err = l.client.DB.Table("links").Delete(model.Link{}, "id = ?", id).Error
	return
}

func (l *LinkService) GetLinksByUserId(id int) (links []model.Link, err error) {
	err = l.client.DB.Table("links").Where("uid = ?", id).Scan(&links).Error
	return
}