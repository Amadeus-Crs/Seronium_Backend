package repository

import (
	"Seronium/internal/config"
	"context"
	"io"

	"github.com/minio/minio-go"
)

func UploadFile(ctx context.Context, filename string, fileData interface{}, fileSize int64, contentType string) (string, error) {
	_, err := MinIOClient.PutObjectWithContext(
		ctx,
		config.MinIOBucket,
		filename,
		fileData.(io.Reader),
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", err
	}
	return filename, nil
}
