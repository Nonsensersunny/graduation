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
	Favorite bool `json:"favorite"`
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

func (c *ContentService) UserGetContentById(id, uid int) (content ReqContent, err error) {
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
	favService := NewFavService(config.GetMySQLClient())
	fav, err := favService.IsFavorite(uid, id)
	if err != nil {
		return ReqContent{}, err
	}
	content.Favorite = fav
	return
}

func (c *ContentService) CreateContent(content model.Content) error {
	if content.Title == "" || content.Content == "" {
		return errors.New("title and content show not be empty")
	}
	return c.client.DB.Table("contents").Create(&content).Error
}

func (c *ContentService) GetRankedContent() (content []model.Content, err error) {
	err = c.client.DB.Table("contents").Find(&content).Order("time desc", true).Error
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
	err = c.client.DB.Table("contents").Where("is_key = ?", true).Scan(&content).Limit(5).Order("views DESC", true).Error
	return
}

func (c *ContentService) GetCategories() (cats []model.Category, err error) {
	err = c.client.DB.Table("categories").Find(&cats).Error
	return
}

func (c *ContentService) CreateCategory(cat model.Category) (err error) {
	err = c.client.DB.Table("categories").Create(&cat).Error
	return
}

func (c *ContentService) ToggleKeyContent(id int) (err error) {
	var content model.Content
	err = c.client.DB.Table("contents").Where("id = ?", id).Scan(&content).Error
	if err != nil {
		return
	}
	err = c.client.DB.Table("contents").Where("id = ?", id).Update("is_key", (!content.IsKey)).Error
	return
}