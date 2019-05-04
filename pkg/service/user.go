package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"graduation/internal/log"
	"graduation/internal/utils"
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
	"strings"
	"time"
)

type UserService struct {
	client *mysql.Client
}

func NewUserService(client *mysql.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) ValidateUser(user model.User) error {
	dbUser, err := s.GetByUsername(user.Username)
	if err != nil {
		return errors.New("User not exist.")
	}

	formPwd := utils.MD5Encode(user.Password, dbUser.Salt)
	if !strings.EqualFold(formPwd, dbUser.Password) {
		return errors.New("Username or password error.")
	}
	return nil
}

func (s *UserService) GetByUsername(username string) (user model.User, err error) {
	err = s.client.DB.Table("users").Where("username=?", username).Scan(&user).Error
	return
}

func (s *UserService) CreateUser(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Username and password should not be empty")
	}

	u, _ := s.GetByUsername(user.Username)
	if u.Username != "" {
		return errors.New("User already exists.")
	}

	user.Salt = utils.RandSalt()
	user.Password = utils.MD5Encode(user.Password, user.Salt)
	return s.client.DB.Table("users").Create(&user).Error
}

func (s *UserService) UpdateUserGrades(id int, cat string) error {
	var category model.Category
	err := s.client.DB.Table("categories").Where("name = ?", cat).Scan(&category).Error
	if err != nil {
		return err
	}
	return s.client.DB.Table("users").Where("id = ?", id).Update("grades", gorm.Expr("grades + ?", category.Weight)).Error
}

type ReqUser struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Grades int `json:"grades"`
	Time time.Time `json:"time"`
	Avatar string `json:"avatar"`
	Description string `json:"description"`
	Mail string `json:"mail"`
	Sex int `json:"sex"`
	Role string `json:"role"`

	ShareNum int `json:"share_num"`
	TopicNum int `json:"topic_num"`
	QuestNum int `json:"quest_num"`
	RecuiNum int `json:"recui_num"`
	CommeNum int `json:"comme_num"`
}

func (s *UserService) GetRankedUsers() (users []ReqUser, err error) {
	err = s.client.DB.Table("users").Find(&users).Order("grades DESC", true).Limit(10).Scan(&users).Error
	return
}

func (s *UserService) GetUserIdByName(username string) (id int, err error) {
	var user model.User
	err = s.client.DB.Table("users").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return
	}
	return user.Id, nil
}

func (s *UserService) GetUserProfileByName(username string) (user ReqUser, err error) {
	id, err := s.GetUserIdByName(username)
	if err != nil {
		return
	}
	return s.GetUserProfileById(id)
}

func (s *UserService) GetUserProfileById(id int) (user ReqUser, err error) {
	err = s.client.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return
	}
	var count int
	s.client.DB.Table("contents").Where("author = ? and category = ?", id, "Share").Count(&count)
	user.ShareNum = count
	s.client.DB.Table("contents").Where("author = ? and category = ?", id, "Topic").Count(&count)
	user.TopicNum= count
	s.client.DB.Table("contents").Where("author = ? and category = ?", id, "Q&A").Count(&count)
	user.QuestNum= count
	s.client.DB.Table("contents").Where("author = ? and category = ?", id, "Recruit").Count(&count)
	user.RecuiNum= count
	s.client.DB.Table("comments").Where("`from` = ?", id).Count(&count)
	user.CommeNum= count
	return
}

func (s *UserService) UpdateUserProfile(user model.User) error {
	log.Infof("%#v", user)
	err := s.client.DB.Table("users").Where("id = ?", user.Id).Updates(&user, true).Error
	log.Infof("%#v", user)
	return err
}

func (s *UserService) GetUsernameById(id int) (string, error) {
	var user model.User
	err := s.client.DB.Table("users").First(&user, id).Error
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
