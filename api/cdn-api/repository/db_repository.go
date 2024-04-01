package repository

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
	"gorm.io/gorm"
)

type DbRepository interface {
	SaveMetadata(request model.Request, env *utils.Environment) (model.Response, error)
}

type DbRepositoryImpl struct {
	db *gorm.DB
}

func NewDbRepository(db *gorm.DB) DbRepository {
	db.AutoMigrate(&model.Metadata{})

	return &DbRepositoryImpl{}
}

func (repo *DbRepositoryImpl) SaveMetadata(request model.Request, env *utils.Environment) (model.Response, error) {

	metadata := model.Metadata{
		Id:       1,
		MimeType: "request.MimeType",
	}
	result := repo.db.Save(&metadata)

	if result.Error != nil {
		return model.Response{}, result.Error
	}

	return model.Response{}, nil
}
