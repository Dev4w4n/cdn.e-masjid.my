package repository

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/enums"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
	"github.com/google/uuid"
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
	response, err := save(request, env)

	if err != nil {
		return model.Response{}, err
	}

	return response, nil
}

func save(request model.Request, env *utils.Environment) (model.Response, error) {

	// check if mime type is allowed
	err := enums.AllowedMimeTypes(request.MimeType)
	if err != nil {
		return model.Response{}, err
	}

	// prepare data
	namespace := request.SubDomain
	fileData := request.Base64File
	fileName := uuid.New().String()
	fileExtension := enums.GetFileExtension(request.MimeType)

	// check if folder exists then cd into it
	err = checkFolderExists(namespace, env)
	if err != nil {
		return model.Response{}, err
	}

	// Create the file
	filePath := filepath.Join(fileName + "." + fileExtension)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		return model.Response{}, err
	}
	defer file.Close()

	// convert fileData base64 to []byte
	var byteData []byte
	byteData, err = base64ToByte(fileData)

	if err != nil {
		log.Printf("Error converting base64 to byte: %v\n", err)
		return model.Response{}, err
	}

	// Write the data to file
	_, err = file.Write(byteData)
	if err != nil {
		log.Printf("Error writing data to file: %v\n", err)
		return model.Response{}, err
	}
	log.Printf("File saved successfully: %s\n", filePath)

	// return response
	response := model.Response{
		Path: namespace + "/" + fileName + "." + fileExtension,
	}

	return response, nil
}

// check if folder for namespace exists
// if not create the namespace folder
func checkFolderExists(namespace string, env *utils.Environment) error {
	repositoryPath := env.RepositoryPath

	// Change directory to the repository path
	if err := os.Chdir(repositoryPath); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return err
	}

	// Create the directory if it doesn't exist
	if _, err := os.Stat(namespace); os.IsNotExist(err) {
		err := os.MkdirAll(namespace, 0755)
		if err != nil {
			log.Printf("Error creating directory: %v\n", err)
			return err
		}
	}

	// Change directory to the namespace folder
	if err := os.Chdir(namespace); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return err
	}

	return nil
}

func base64ToByte(base64Data string) ([]byte, error) {
	// Decode the base64 encoded data
	byteData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}
