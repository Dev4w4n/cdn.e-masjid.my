package service

import (
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/repository"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/utils"
)

type UploadService interface {
	SaveFile(request model.Request, env *utils.Environment) (model.Response, error)
}

type UploadServiceImpl struct {
	dbRepository   repository.DbRepository
	fileRepository repository.FileRepository
}

func NewUploadService(dbRepository repository.DbRepository, fileRepository repository.FileRepository) UploadService {
	return &UploadServiceImpl{
		dbRepository:   dbRepository,
		fileRepository: fileRepository,
	}
}

func (repo *UploadServiceImpl) SaveFile(request model.Request, env *utils.Environment) (model.Response, error) {

	repo.dbRepository.SaveFile(request, env)
	repo.fileRepository.SaveImage(request, env)
	return model.Response{}, nil
}
