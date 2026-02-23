package service

import (
	"Seronium/internal/config"
	"Seronium/internal/repository"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type UploadService struct{}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (s *UploadService) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%d/%s%s", time.Now().Year(), uuid.New().String(), ext)

	key, err := repository.UploadFile(repository.CTX, filename, file, fileHeader.Size, fileHeader.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://%s/%s/%s",
		config.MinIOEndpoint,
		config.MinIOBucket,
		key,
	)

	return url, nil
}
