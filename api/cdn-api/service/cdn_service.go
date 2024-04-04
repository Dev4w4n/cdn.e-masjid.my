package service

import (
	"os"

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
	response, err := repo.fileRepository.SaveFile(request, env)

	if err != nil {
		return model.Response{}, err
	}

	endpoint := "https://cdn.e-masjid.my/volume"
	isLocalEnv := os.Getenv("GO_ENV")
	if isLocalEnv == "" || isLocalEnv == "dev" {
		endpoint = "http://localhost/volume"
	}

	metadata := model.Metadata{
		MimeType:       request.MimeType,
		SubDomain:      request.SubDomain,
		TableReference: request.TableReference,
		MarkAsDelete:   request.MarkAsDelete,
		Path:           endpoint + "/" + response.Path,
	}

	dbResponse, err := repo.dbRepository.SaveMetadata(metadata)

	if err != nil {
		return model.Response{}, err
	}

	return dbResponse, nil
}
