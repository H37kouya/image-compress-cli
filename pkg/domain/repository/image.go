package repository

import "github.com/H37kouya/image-compress-cli/pkg/domain/model"

// ImageRepository Image に関するリポジトリ
type ImageRepository interface {
	ResizeImage(file model.File, newfiledir, newfilename string, width, quality int) (model.File, error)
	ResizeJpg(originalFile model.File, newfiledir, newfilename string, width, quality int) (model.File, error)
}
