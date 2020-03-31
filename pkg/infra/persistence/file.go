package persistence

import (
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/domain/repository"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FilePersistence File データの構造体
type filePersistence struct{}

// NewFilePersistence : File データに関する Persistence を生成
func NewFilePersistence() repository.FileRepository {
	return &filePersistence{}
}

// GetFileListsFromStorages ファイルの一覧を作成
func (fp filePersistence) GetFileListsFromStorages(dir string) model.Files {
	// file一覧を作成
	fileinfos, _ := ioutil.ReadDir(dir)
	// model.Files型に詰める
	files := make([]model.File, 0)
	for _, f := range fileinfos {
		file := convertToFile(dir, f)
		files = append(files, file)
	}

	return files
}

func convertToFile(dir string, f os.FileInfo) model.File {
	// ファイル名
	filename := f.Name()
	// 拡張子
	ext := strings.TrimLeft(filepath.Ext(filename), ".")
	extLowerCase := strings.ToLower(ext)
	name := getFileNameWithoutExt(filename)
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
