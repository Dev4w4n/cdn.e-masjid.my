package service

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/repository"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
)

type CDNService interface {
	Upload(request model.Request, env *utils.Environment) (model.Response, error)
}

type CDNServiceImpl struct {
	dbRepository   repository.DbRepository
	fileRepository repository.FileRepository
}

func NewCDNService(dbRepository repository.DbRepository, fileRepository repository.FileRepository) CDNService {
	return &CDNServiceImpl{
		dbRepository:   dbRepository,
		fileRepository: fileRepository,
	}
}

func (repo *CDNServiceImpl) Upload(request model.Request, env *utils.Environment) (model.Response, error) {
	repo.fileRepository.SaveFile(request, env)

	err := repo.dbRepository.SaveMetadata(request)

	if err != nil {
		return model.Response{}, err
	}

	return model.Response{}, nil
}
