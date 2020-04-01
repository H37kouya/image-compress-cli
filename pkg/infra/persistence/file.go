package persistence

import (
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/domain/repository"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// NewFilePersistence : File データに関する Persistence を生成
func NewFilePersistence() repository.FileRepository {
	return &filePersistence{}
}

// filePersistence File データの構造体
type filePersistence struct{}

// GetFileListsFromStorages ファイルの一覧を作成
func (fp filePersistence) GetFileListsFromStorages(dir string) (model.Files, error) {
	// file一覧を作成
	fileinfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// model.Files型に詰める
	files := make([]model.File, 0)
	for _, f := range fileinfos {
		file := convertToFile(dir, f)
		files = append(files, file)
	}

	return files, nil
}

func convertToFile(dir string, f os.FileInfo) model.File {
	// ファイル名
	filename := f.Name()
	// 拡張子
	ext := strings.TrimLeft(filepath.Ext(filename), ".")
	extLowerCase := strings.ToLower(ext)
	name := getFileNameWithoutExt(filename)
	if name == "." { // .gitkeepのようなファイルを考慮
		name = ""
	}
	filename = name + "." + extLowerCase
	path := strings.TrimRight(dir, "/") + "/" + filename

	return model.File{
		Extension:    ext,
		ExtLowerCase: extLowerCase,
		FileName:     filename,
		Name:         name,
		Path:         path,
	}
}

// getFileNameWithoutExt 拡張子なしのファイル名を取得する
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
