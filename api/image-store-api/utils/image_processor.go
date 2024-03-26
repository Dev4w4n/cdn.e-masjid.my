package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/model"
)

func InitializeImageFolders(env *Environment) error {
	namespace := env.Namespace
	workdir := env.Workdir
	repositoryPath := env.RepositoryPath

	// Change directory to the repositoryPath path
	if err := os.Chdir(repositoryPath); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return err
	}

	// Create the images folder if it doesn't exist
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		err := os.Mkdir("images", 0755)
		if err != nil {
			log.Printf("Error creating images folder: %v\n", err)
			return err
		}
		log.Printf("Images folder created successfully.\n")
	}

	// Copy Workdir main.jpg to repositoryPath images folder
	workFile := filepath.Join(workdir, "main.jpg")
	repositoryFile := filepath.Join(repositoryPath, "images", "main.jpg")
	log.Printf("Copying " + workFile + " to " + repositoryFile + ".\n")
	err := copyFile(workFile, repositoryFile)
	if err != nil {
		log.Printf("Error copying file: %v\n", err)
		return err
	}
	log.Printf("Successfull.\n")

	// Change directory to the repository path
	if err := os.Chdir(repositoryPath + "/images"); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return err
	}

	// Split the namespace by comma to get individual folder names
	folders := strings.Split(namespace, ",")

	// Copy main.jpg file to each newly created folder
	sourceFile := filepath.Join(repositoryPath, "images", "main.jpg") // Change this to the actual file name you want to copy
	// Iterate through each folder name and check if it exists
	for _, folder := range folders {
		// Check if the folder exists, if not create it
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err := os.Mkdir(folder, 0755)
			if err != nil {
				log.Printf("Error creating folder %s: %v\n", folder, err)
				return err
			}
			log.Printf("Folder %s created successfully.\n", folder)
		}

		// Copy the file into the folder
		destination := filepath.Join(repositoryPath, "images", folder, "main.jpg")
		if _, err := os.Stat(destination); os.IsNotExist(err) {
			err := copyFile(sourceFile, destination)
			if err != nil {
				log.Printf("Error copying file %s to folder %s: %v\n", sourceFile, destination, err)
			} else {
				log.Printf("File %s copied to folder %s.\n", sourceFile, destination)
			}
		}
	}
	return nil
}

// CopyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
func SaveImage(request model.Request, env *Environment) (model.Response, error) {
	response := model.Response{}
	switch request.ImageType {
	case 1:
		res, err := saveMasjidImage(request, env)
		if err != nil {
			return model.Response{}, err
		}
		response = res
	default:
		return model.Response{}, fmt.Errorf("unsupported image type")
	}
	return response, nil
}

func saveMasjidImage(request model.Request, env *Environment) (model.Response, error) {
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

// func savePublicImage(request model.Request, env *Environment) error {
// 	return nil
// }
