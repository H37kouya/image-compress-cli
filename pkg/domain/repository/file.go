package repository

import "image-compress-cli/pkg/domain/model"

// FileRepository ファイルリポジトリ
type FileRepository interface {
	GetFileListsFromStorages(dir string) (model.Files, error)
}
