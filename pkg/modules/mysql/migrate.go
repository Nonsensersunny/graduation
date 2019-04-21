package mysql

import "graduation/pkg/modules/model"

func MigrateTables(client *Client) {
	client.DB.Debug().AutoMigrate(model.User{})
}