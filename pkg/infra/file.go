package infra

import (
	"os"
	"uu-circle-pic/pkg/domain/model"
)

// CreateNewFile ファイルの新規作成
func CreateNewFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

// GetFileListsFromStorages ファイルの一覧を作成
func GetFileListsFromStorages(dir string) model.Files {
	// path := "./" + strings.TrimLeft(dir, "/")

	// TODO: file一覧を作成
	// TODO: model.Files型に詰める
	files := make([]model.File, 0)
	file := model.File{
		Name:      "example",
		Extension: "jpg",
	}

	files = append(files, file)

	return files
}
