package services

import (
	"os"
)

type FileService interface {
	ReadFile(day string, path string) (string, error)
}

type fileService struct{}

func NewFileService() FileService {
	return &fileService{}
}

/**
 * ReadFile reads the file at the given path and returns the bytes
 */
func (fileService *fileService) ReadFile(day string, filename string) (string, error) {

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(pwd + "/" + day + "/" + filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
