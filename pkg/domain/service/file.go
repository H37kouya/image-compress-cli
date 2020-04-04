package service

import (
	"image-compress-cli/pkg/domain/repository"
)

type fileService struct {
	fileRepository repository.FileRepository
}

// FileService FileServiceのためのinterface
type FileService interface {
}

// NewFileService FileService DIするために必要
func NewFileService(fr repository.FileRepository) FileService {
	return &fileService{
		fileRepository: fr,
	}
}
