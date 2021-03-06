package di

import (
	"github.com/H37kouya/image-compress-cli/pkg/usecase"
)

// InjectImageUseCase ImageUseCaseの依存性注入
func InjectImageUseCase() usecase.ImageUseCase {
	return usecase.NewImageUseCase(
		InjectFilePersistence(),
		InjectImagePersistence(),
		InjectImageConfig(),
	)
}
