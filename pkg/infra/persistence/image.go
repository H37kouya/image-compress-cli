package persistence

import (
	"image"
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/domain/repository"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

// NewImagePersistence : Image データに関する Persistence を生成
func NewImagePersistence() repository.ImageRepository {
	return &imagePersistence{}
}

// imagePersistence Image データの構造体
type imagePersistence struct{}

// ResizeImage 画像をリサイズする
func (ip imagePersistence) ResizeImage(file model.File, newfiledir, newfilename string, width, quality int) (model.File, error) {
	if file.ExtLowerCase == "jpg" || file.ExtLowerCase == "jpeg" {
		newfile, err := ip.ResizeJpg(file, newfiledir, newfilename, width, quality)
		if err != nil {
			return model.File{}, err
		}
		return newfile, nil
	}
	return file, nil
}

// ResizeJpg jpegをリサイズし新規ファイルに保存
func (ip imagePersistence) ResizeJpg(originalFile model.File, newfiledir, newfilename string, width, quality int) (model.File, error) {
	srcImage, err := getJpegFromStorages(originalFile.Path)
	if err != nil {
		return model.File{}, err
	}

	// ディレクトリがない場合は新規作成
	if _, err := os.Stat(newfiledir); os.IsNotExist(err) {
		os.Mkdir(newfiledir, 0777)
	}

	dir := strings.TrimRight(newfiledir, "/") + "/"
	path := dir + newfilename

	// ファイルの新規作成
	f, err := os.Create(path)
	if err != nil {
		return model.File{}, err
	}
	defer f.Close()

	// リサイズする
	resizeImage := imaging.Resize(srcImage, width, 0, imaging.Linear)

	// リサイズした画像を新規ファイルに保存
	if err := imaging.Save(resizeImage, f.Name(), imaging.JPEGQuality(quality)); err != nil {
		return model.File{}, err
	}

	return model.File{
		Dir:          newfiledir,
		Extension:    originalFile.Extension,
		ExtLowerCase: originalFile.ExtLowerCase,
		FileName:     newfilename,
		Name:         getFileNameWithoutExt(newfilename),
		Path:         path,
	}, nil
}

// getJpegFromStorages ファイルをstorageから取得する
func getJpegFromStorages(path string) (image.Image, error) {
	// 画像の取得
	srcImage, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}

	return srcImage, nil
}
