package repository

import "github.com/H37kouya/image-compress-cli/pkg/domain/model"

// FileRepository ファイルリポジトリ
type FileRepository interface {
	GetFileListsFromStorages(dir string) (model.Files, error)
}
