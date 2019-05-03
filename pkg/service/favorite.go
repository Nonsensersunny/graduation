package service

import (
	"graduation/pkg/modules/model"
	"graduation/pkg/modules/mysql"
)

type FavoriteService struct {
	client *mysql.Client
}

func NewFavService(client *mysql.Client) *FavoriteService {
	return &FavoriteService{
		client: client,
	}
}

func (f *FavoriteService) CreateFavorite(uid, cid int) (err error) {
	fav := model.Favorite{
		Uid: uid,
		Cid: cid,
		Valid: true,
	}
	err = f.client.DB.Table("favorites").Create(&fav).Error
	return
}

func (f *FavoriteService) DeleteFavorite(uid, cid int) (err error) {
	err = f.client.DB.Table("favorites").Where("uid = ? and cid = ?", uid, cid).Set("valid", false).Error
	if err != nil {
		return f.client.DB.Table("favorites").Create(&model.Favorite{
			Cid: cid,
			Uid: uid,
			Valid: false,
		}).Error
	}
	return
}

func (f *FavoriteService) IsFavorite(uid, cid int) (valid bool, err error) {
	var fav model.Favorite
	err = f.client.DB.Table("favorites").Where("uid = ? and cid = ?", uid, cid).Scan(&fav).Error
	if err != nil {
		return false, f.DeleteFavorite(uid, cid)
	}
	return fav.Valid, nil
}

