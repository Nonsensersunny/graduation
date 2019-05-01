package mysql

import "graduation/pkg/modules/model"

func MigrateTables(client *Client) {
	client.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8")
	client.DB.Set("character", "utf8")
	if err := client.DB.Debug().AutoMigrate(model.User{}, model.Content{}, model.Comment{}, model.Favorite{}, model.Category{}).Error; err != nil {
		panic(err)
	}
}