package di

import (
	"github.com/H37kouya/image-compress-cli/pkg/domain/service"
)

// InjectFileService InjectFileServiceの依存性注入
func InjectFileService() service.FileService {
	return service.NewFileService(
		InjectFilePersistence(),
	)
}
