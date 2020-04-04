package di

import (
	"image-compress-cli/pkg/domain/service"
)

// InjectFileService InjectFileServiceの依存性注入
func InjectFileService() service.FileService {
	return service.NewFileService(
		InjectFilePersistence(),
	)
}
