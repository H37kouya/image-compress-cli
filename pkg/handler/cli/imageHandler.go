package cli

import (
	"github.com/H37kouya/image-compress-cli/pkg/usecase"

	"github.com/spf13/cobra"
)

// ImageHandler : Image における Handler のインターフェース
type ImageHandler interface {
	Index(cmd *cobra.Command)
}

type imageHandler struct {
	imageUseCase usecase.ImageUseCase
}

// NewImageHandler : Image データに関する Handler を生成
func NewImageHandler(iu usecase.ImageUseCase) ImageHandler {
	return &imageHandler{
		imageUseCase: iu,
	}
}

// Index index
func (ih imageHandler) Index(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	width, err := flags.GetInt("width")
	if err != nil {
		panic(err)
	}

	ih.imageUseCase.ResizeImages(width)
}
