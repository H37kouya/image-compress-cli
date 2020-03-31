package usecase

import (
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/domain/repository"
)

// ImageUseCase Imageに関するUseCase
type ImageUseCase interface {
	ResizeImages(width int) model.Files
}

type imageUseCase struct {
	fileRepository  repository.FileRepository
	imageRepository repository.ImageRepository
}

// NewImageUseCase : Image データに関する UseCase を生成
func NewImageUseCase(fr repository.FileRepository, ir repository.ImageRepository) ImageUseCase {
	return &imageUseCase{
		fileRepository:  fr,
		imageRepository: ir,
	}
}

// ResizeImages 画像のリサイズをする
func (iu imageUseCase) ResizeImages(width int) model.Files {
	sizes := make([]int, 0, 7)
	if width == 0 {
		sizes = append(sizes, width)
	} else {
		sizes = append(sizes, 80, 320, 480, 640, 960, 1200, 1800)
	}

	files := iu.fileRepository.GetFileListsFromStorages("./storages/before")

	for _, file := range files {

		for _, size := range sizes {
			iu.imageRepository.ResizeImage(file, "./storages/after/"+file.FileName, size, 70)
		}
	}

	return files
}
