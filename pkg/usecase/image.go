package usecase

import (
	"strconv"

	"github.com/H37kouya/image-compress-cli/pkg/config"
	"github.com/H37kouya/image-compress-cli/pkg/domain/model"
	"github.com/H37kouya/image-compress-cli/pkg/domain/repository"
)

// ImageUseCase Imageに関するUseCase
type ImageUseCase interface {
	ResizeImages(width int) (model.Files, error)
}

type imageUseCase struct {
	fileRepository  repository.FileRepository
	imageRepository repository.ImageRepository
	imageConfig     config.ImageConfig
}

// NewImageUseCase : Image データに関する UseCase を生成
func NewImageUseCase(fr repository.FileRepository, ir repository.ImageRepository, ic config.ImageConfig) ImageUseCase {
	return &imageUseCase{
		fileRepository:  fr,
		imageRepository: ir,
		imageConfig:     ic,
	}
}

// ResizeImages 画像のリサイズをする
func (iu imageUseCase) ResizeImages(width int) (model.Files, error) {
	config := iu.imageConfig.LoadImageConfig()
	conSizes := config.Images.Sizes
	sizes := make([]int, 0, len(conSizes))

	if width != 0 {
		sizes = append(sizes, width)
	} else {
		for _, size := range conSizes {
			sizes = append(sizes, size)
		}
	}

	files, err := iu.fileRepository.GetFileListsFromStorages("./storages/before")
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		for _, size := range sizes {
			iu.imageRepository.ResizeImage(file, "./storages/after/", strconv.Itoa(size)+"-"+file.FileName, size, 70)
		}
	}

	return files, nil
}
