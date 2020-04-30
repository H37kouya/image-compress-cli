package di

import "github.com/H37kouya/image-compress-cli/pkg/config"

// InjectImageConfig ImageConfigの依存性注入
func InjectImageConfig() config.ImageConfig {
	return config.NewImageConfig()
}
