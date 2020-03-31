package di

import (
	"image-compress-cli/pkg/handler/cli"
)

// InjectImageHandler ImageHandlerの依存性注入
func InjectImageHandler() cli.ImageHandler {
	return cli.NewImageHandler(
		InjectImageUseCase(),
	)
}
