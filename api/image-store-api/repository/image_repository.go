package repository

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/utils"
)

type ImageRepository interface {
	SaveImage(request model.Request, env *utils.Environment) (model.Response, error)
}

type ImageRepositoryImpl struct {
}

func NewImageRepository() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (i *ImageRepositoryImpl) SaveImage(request model.Request, env *utils.Environment) (model.Response, error) {
	response, err := utils.SaveImage(request, env)

	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}
