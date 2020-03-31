package persistence

import (
	"image"
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/domain/repository"
	"os"

	"github.com/disintegration/imaging"
)

// ImagePersistence Image データの構造体
type imagePersistence struct{}

// NewImagePersistence : Image データに関する Persistence を生成
func NewImagePersistence() repository.ImageRepository {
	return &imagePersistence{}
}

// ResizeImage 画像をリサイズする
func (ip imagePersistence) ResizeImage(file model.File, newfilepath string, width, quality int) model.File {
	if file.ExtLowerCase == "jpg" || file.ExtLowerCase == "jpeg" {
		ip.ResizeJepgToStorages(file.Path, newfilepath, width, quality)
		return file
	}
	return file
}

// ResizeJepgToStorages jpegをリサイズし新規ファイルに保存
func (ip imagePersistence) ResizeJepgToStorages(originalfilepath, newfilepath string, width, quality int) {
	srcImage := getJpegFromStorages(originalfilepath)

	// ファイルの新規作成
	f, err := os.Create(newfilepath)
	if err != nil {
		panic(err)
	}

	// リサイズする
	resizeImage := imaging.Resize(srcImage, width, 0, imaging.Linear)

	// リサイズした画像を新規ファイルに保存
	if err := imaging.Save(resizeImage, f.Name(), imaging.JPEGQuality(quality)); err != nil {
		panic(err)
	}
	f.Close()
}

// getJpegFromStorages ファイルをstorageから取得する
func getJpegFromStorages(path string) image.Image {
	// 画像の取得
	srcImage, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}

	return srcImage
}
