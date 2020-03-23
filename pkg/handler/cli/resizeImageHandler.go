package cli

import (
	"fmt"
	"image-compress-cli/pkg/infra"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ResizeImageHandler 画像のリサイズをする
func ResizeImageHandler(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	width, err := flags.GetInt("width")
	if err != nil {
		panic(err)
	}

	// 色指定
	green := color.New(color.FgGreen).SprintFunc()
	// TODO: ディレクトリから画像の一覧作成
	dirLists := infra.GetFileListsFromStorages("./storages/before")

	for _, file := range dirLists {
		// 画像の取得
		img := infra.GetJpegFromStorages("/before/" + file.Name)
		// 画像のリサイズ
		infra.ResizeJepgToStorages(img, "/after/"+file.Name, width, 70)

		log.Println(green("Succeess: ") + file.Name)
	}

	fmt.Println("")
	log.Println(green("All Succeess"))
}
