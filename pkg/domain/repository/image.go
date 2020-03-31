package repository

import "image-compress-cli/pkg/domain/model"

// ImageRepository Image に関するリポジトリ
type ImageRepository interface {
	ResizeImage(file model.File, newfilepath string, width, quality int) model.File
}
