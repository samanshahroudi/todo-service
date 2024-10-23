package ports

import "io"

type S3Service interface {
	UploadFile(key string, file io.Reader) (string, error)
}
