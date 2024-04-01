package repository

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/utils"
)

type DbRepository interface {
	SaveFile(request model.Request, env *utils.Environment) (model.Response, error)
}

type DbRepositoryImpl struct {
}

func NewDbRepository() DbRepository {
	return &DbRepositoryImpl{}
}

func (i *DbRepositoryImpl) SaveFile(request model.Request, env *utils.Environment) (model.Response, error) {

	return model.Response{}, nil
}
