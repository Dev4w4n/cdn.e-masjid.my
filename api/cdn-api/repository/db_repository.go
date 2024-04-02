package repository

import (
	"time"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"gorm.io/gorm"
)

type DbRepository interface {
	SaveMetadata(metadata model.Metadata) error
}

type DbRepositoryImpl struct {
	Db *gorm.DB
}

func NewDbRepository(db *gorm.DB) DbRepository {
	db.AutoMigrate(&model.Metadata{})

	return &DbRepositoryImpl{Db: db}
}

func (repo *DbRepositoryImpl) SaveMetadata(metadata model.Metadata) error {

	tuple := metadata
	tuple.CreateDate = time.Now().AddDate(time.Now().Year(), 1, 1).Unix() / 1000
	result := repo.Db.Save(&tuple)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
