package persistence_test

import (
	"testing"

	"github.com/H37kouya/image-compress-cli/pkg/domain/model"
	"github.com/H37kouya/image-compress-cli/pkg/infra/persistence"
)

func TestResizeImageSuccess(t *testing.T) {
	ip := persistence.NewImagePersistence()
	var originalFile model.File

	t.Run("Jpgのリサイズをする", func(t *testing.T) {
		originalFile = model.File{
			Dir:          "../../../tests/storages/before/",
			Extension:    "jpg",
			ExtLowerCase: "jpg",
			FileName:     "example1.jpg",
			Name:         "example1",
			Path:         "../../../tests/storages/before/example1.jpg",
		}
		resultFile, err := ip.ResizeImage(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err != nil {
			t.Fatalf("\nfailed test %#v", err)
		}
		if resultFile.ExtLowerCase != originalFile.ExtLowerCase {
			t.Errorf("\nfailed test \n実際：%#v \n期待：%#v", resultFile.ExtLowerCase, originalFile.ExtLowerCase)
		}
	})
}

// TestResizeJpgSuccess ResizeJpgの正常系テスト
func TestResizeJpgSuccess(t *testing.T) {
	ip := persistence.NewImagePersistence()
	originalFile := model.File{
		Dir:          "../../../tests/storages/before/",
		Extension:    "jpg",
		ExtLowerCase: "jpg",
		FileName:     "example1.jpg",
		Name:         "example1",
		Path:         "../../../tests/storages/before/example1.jpg",
	}

	t.Run("Jpgのリサイズをする", func(t *testing.T) {
		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err != nil {
			t.Fatalf("\nfailed test %#v", err)
		}
	})

	t.Run("JPGのリサイズをする", func(t *testing.T) {
		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example2.jpg", 960, 70,
		)
		if err != nil {
			t.Fatalf("\nfailed test %#v", err)
		}
	})
}

// TestResizeJpgFailed ResizeJpgの異常系テスト
func TestResizeJpgFailed(t *testing.T) {
	ip := persistence.NewImagePersistence()

	t.Run("originalpathでファイルが存在しない", func(t *testing.T) {
		originalFile := model.File{
			Dir:          "../../../tests/storages/before",
			Extension:    "",
			ExtLowerCase: "",
			FileName:     "",
			Name:         "",
			Path:         "../../../tests/storages/before",
		}

		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err == nil {
			t.Fatalf("\nfailed test")
		}
	})

	t.Run("originalpathでファイルが画像でない", func(t *testing.T) {
		originalFile := model.File{
			Dir:          "../../../tests/storages/before/",
			Extension:    "gitkeep",
			ExtLowerCase: "gitkeep",
			FileName:     ".gitkeep",
			Name:         "",
			Path:         "../../../tests/storages/before/.gitkeep",
		}

		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err == nil {
			t.Fatalf("\nfailed test")
		}
	})
}
