package repository

import (
	"time"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"gorm.io/gorm"
)

type DbRepository interface {
	SaveMetadata(metadata model.Metadata) (model.Response, error)
}

type DbRepositoryImpl struct {
	Db *gorm.DB
}

func NewDbRepository(db *gorm.DB) DbRepository {
	db.AutoMigrate(&model.Metadata{})

	return &DbRepositoryImpl{Db: db}
}

func (repo *DbRepositoryImpl) SaveMetadata(metadata model.Metadata) (model.Response, error) {

	// Set the CreateDate field of the metadata
	metadata.CreateDate = time.Now().AddDate(time.Now().Year(), 1, 1).Unix() / 1000

	// Save the metadata to the database
	result := repo.Db.Save(&metadata)
	if result.Error != nil {
		return model.Response{}, result.Error
	}

	// Prepare the response with the persisted metadata details
	response := model.Response{
		ID:        metadata.Id,
		Path:      metadata.Path,
		CreatedAt: metadata.CreateDate,
	}

	return response, nil
}
