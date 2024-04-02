package repository

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"gorm.io/gorm"
)

type DbRepository interface {
	SaveMetadata(request model.Request) error
}

type DbRepositoryImpl struct {
	db *gorm.DB
}

func NewDbRepository(db *gorm.DB) DbRepository {
	db.AutoMigrate(&model.Metadata{})

	return &DbRepositoryImpl{}
}

func (repo *DbRepositoryImpl) SaveMetadata(request model.Request) error {

	metadata := model.Metadata{
		TableReference: request.TableReference,
		SubDomain:      request.SubDomain,
		MarkAsDelete:   request.MarkAsDelete,
		MimeType:       request.MimeType,
	}
	result := repo.db.Save(&metadata)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
