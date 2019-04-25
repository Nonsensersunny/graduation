package service

import (
	"errors"
	"graduation/internal/utils"
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
	"strings"
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