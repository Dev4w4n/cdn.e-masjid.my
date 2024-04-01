package repository

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
)

type DbRepository interface {
	SaveMetadata(request model.Request, env *utils.Environment) (model.Response, error)
}

type DbRepositoryImpl struct {
}

func NewDbRepository() DbRepository {
	return &DbRepositoryImpl{}
}

func (i *DbRepositoryImpl) SaveMetadata(request model.Request, env *utils.Environment) (model.Response, error) {

	return model.Response{}, nil
}
