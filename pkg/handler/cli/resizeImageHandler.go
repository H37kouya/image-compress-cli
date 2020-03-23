package cli

import (
	"fmt"
	"image-compress-cli/pkg/infra"
)

// ResizeImageHandler 画像のリサイズをする
func ResizeImageHandler() {
	fmt.Println("resizeImage called")

	// TODO: ディレクトリから画像の一覧作成
	dirLists := infra.GetFileListsFromStorages("./storages/before")
	fmt.Println(dirLists)

	// img := infra.GetJpegFromStorages("/before/example.jpg")
	// infra.ResizeJepgToStorages(img, "/after/CatmullRom70.jpg", 960, 70)
	// infra.ResizeJepgToStorages(img, "/after/CatmullRom80.jpg", 960, 80)
	fmt.Println("succeess")
}
