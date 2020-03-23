package infra

import (
	"image"
	"strings"

	"github.com/disintegration/imaging"
)

// GetJpegFromStorages ファイルをstorageから取得する
func GetJpegFromStorages(filename string) image.Image {

	path := "./storages/" + filename

	// 画像の取得
	srcImage, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}

	return srcImage
}

// StoreJpegToStorages jpegを保存
// func StoreJpegToStorages() {

// }

// ResizeJepgToStorages jpegをリサイズし新規ファイルに保存
func ResizeJepgToStorages(srcImage image.Image, filename string, width, quality int) {
	path := "./storages/" + strings.TrimLeft(filename, "/")

	// ファイルの新規作成
	f := CreateNewFile(path)

	// リサイズする
	resizeImage := imaging.Resize(srcImage, width, 0, imaging.Linear)

	// リサイズした画像を新規ファイルに保存
	if err := imaging.Save(resizeImage, f.Name(), imaging.JPEGQuality(quality)); err != nil {
		panic(err)
	}
	f.Close()
}
