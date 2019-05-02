package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"graduation/internal/config"
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

type ReqContent struct {
	Content model.Content `json:"content"`
	Writer ReqUser `json:"writer"`
	Comments []model.Comment `json:"comments"`
}

func (c *ContentService) GetContentById(id int) (content ReqContent, err error) {
	err = c.client.DB.Table("contents").Where("id = ?", id).Scan(&content.Content).Error
	if err != nil {
		return ReqContent{}, err
	}
	userService := NewUserService(config.GetMySQLClient())
	writer, err := userService.GetUserProfileById(content.Content.Author)
	if err != nil {
		return ReqContent{}, err
	}
	content.Writer = writer
	commentService := NewCommentService(config.GetMySQLClient())
	commments, err := commentService.GetCommentsByCid(content.Content.Id)
	if err != nil {
		return ReqContent{}, err
	}
	content.Comments = commments
	return
}

func (c *ContentService) CreateContent(content model.Content) error {
	if content.Title == "" || content.Content == "" {
		return errors.New("title and content show not be empty")
	}
	return c.client.DB.Table("contents").Create(&content).Error
}

func (c *ContentService) GetRankedContent() (content []model.Content, err error) {
	err = c.client.DB.Table("contents").Find(&content).Order("views DESC", true).Error
	return
}

func (c *ContentService) GetContentByCat(cat string) (content []model.Content, err error) {
	err = c.client.DB.Table("contents").Where("category = ?", cat).Find(&content).Error
	return
}

func (c *ContentService) ContentVisited(id int) error {
	return c.client.DB.Table("contents").Where("id = ?", id).Update("views", gorm.Expr("views + ?", 1)).Error
}

func (c *ContentService) GetTopContent() (content []model.Content, err error) {
	err = c.client.DB.Table("contents").Find(&content).Limit(5).Order("views DESC", true).Error
	return
}

func (c *ContentService) GetCategories() (cats []model.Category, err error) {
	err = c.client.DB.Table("categories").Find(&cats).Error
	return
}