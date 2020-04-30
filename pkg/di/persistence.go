package di

import (
	"github.com/H37kouya/image-compress-cli/pkg/domain/repository"
	"github.com/H37kouya/image-compress-cli/pkg/infra/persistence"
)

// InjectFilePersistence FilePersistenceの依存性注入
func InjectFilePersistence() repository.FileRepository {
	return persistence.NewFilePersistence()
}

// InjectImagePersistence ImagePersistenceの依存性注入
func InjectImagePersistence() repository.ImageRepository {
	return persistence.NewImagePersistence()
}
