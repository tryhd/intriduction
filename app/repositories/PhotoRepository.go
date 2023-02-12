package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	InsertPhoto(photo string, id int) models.Profile
	GetPhoto(id int) (res string)
	DeletePhoto(id int) (res models.Profile)
}

type photoConnection struct{ connection *gorm.DB }

func NewPhotoRepository(dbConn *gorm.DB) PhotoRepository { return &photoConnection{connection: dbConn} }

func (db *photoConnection) InsertPhoto(photo string, id int) models.Profile {
	res := models.Profile{}
	db.connection.Model(&res).Where("profile_code", id).Update("photo_url", photo)
	db.connection.Select("profile_code", "photo_url").Where("profile_code", id).Find(&res)
	return res
}

func (db *photoConnection) GetPhoto(id int) (res string) {
	profile := models.Profile{}
	db.connection.Select("photo_url").Where("profile_code", id).Find(&profile)
	return profile.PhotoUrl
}

func (db *photoConnection) DeletePhoto(id int) (res models.Profile) {
	profile := models.Profile{}
	db.connection.Debug().Model(&res).Where("profile_code", id).Update("photo_url", "")
	db.connection.Select("profile_code").Where("profile_code", id).Find(&profile)
	return profile
}
