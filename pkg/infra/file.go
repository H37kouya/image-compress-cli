package infra

import (
	"image-compress-cli/pkg/domain/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	path := "./" + strings.TrimLeft(dir, "/")

	// file一覧を作成
	fileinfos, _ := ioutil.ReadDir(path)
	// model.Files型に詰める
	files := make([]model.File, 0)
	for _, f := range fileinfos {
		fileName := f.Name()
		ext := strings.TrimLeft(filepath.Ext(fileName), ".")

		file := model.File{
			Name:      fileName,
			Extension: ext,
		}
		files = append(files, file)
	}

	return files
}
