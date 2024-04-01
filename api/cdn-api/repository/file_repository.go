package repository

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
)

type FileRepository interface {
	SaveFile(request model.Request, env *utils.Environment) (model.Response, error)
}

type FileRepositoryImpl struct {
}

func NewFileRepository() FileRepository {
	return &FileRepositoryImpl{}
}

func (i *FileRepositoryImpl) SaveFile(request model.Request, env *utils.Environment) (model.Response, error) {
	response, err := saveMasjidImage(request, env)

	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}

func saveMasjidImage(request model.Request, env *utils.Environment) (model.Response, error) {
	namespace := request.Namespace
	imageData := request.Data
	dataType := request.DataType
	imageName := "main"

	repositoryPath := env.RepositoryPath
	folderPath := filepath.Join(repositoryPath, namespace)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Printf("Error creating directory: %v\n", err)
		return model.Response{}, err
	}

	// Change directory to the repository path
	if err := os.Chdir(folderPath); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return model.Response{}, err
	}

	// Write the image data to file
	filePath := filepath.Join(folderPath, imageName+"."+dataType)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		return model.Response{}, err
	}
	defer file.Close()

	// Write image data to the file
	_, err = file.Write(imageData)
	if err != nil {
		log.Printf("Error writing image data to file: %v\n", err)
		return model.Response{}, err
	}

	log.Printf("Image saved successfully: %s\n", filePath)

	response := model.Response{
		Message: imageName + "." + dataType,
	}
	return response, nil
}
