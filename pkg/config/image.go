package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// NewImageConfig : Image データに関する Config を生成
func NewImageConfig() ImageConfig {
	return &imageConfig{}
}

// ImageConfig Imageに関するinterface
type ImageConfig interface {
	LoadImageConfig() Config
}

// imageConfig Image データのConfig
type imageConfig struct {
}

// Config config構造体
type Config struct {
	Images Image `yaml:"image"`
}

// Image image構造体
type Image struct {
	Sizes []int `yaml:"sizes"`
}

// LoadImageConfig ImageConfigをロードする
func (ic imageConfig) LoadImageConfig() Config {
	config, err := getYaml("./config/image.yaml")
	if err != nil {
		panic(err)
	}
	return config
}

func getYaml(path string) (Config, error) {
	// yamlを読み込む
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	// structにUnmasrshal
	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
