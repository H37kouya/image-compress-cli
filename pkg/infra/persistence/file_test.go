package persistence_test

import (
	"image-compress-cli/pkg/di"
	"image-compress-cli/pkg/domain/model"
	"testing"
)

// TestGetFileListsFromStorages GetFileListsFromStoragesの正常系テスト
func TestGetFileListsFromStoragesSuccess(t *testing.T) {
	fp := di.InjectFilePersistence()
	result, err := fp.GetFileListsFromStorages("../../../tests/storages/before")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	expect := model.Files{
		model.File{
			Extension:    "gitkeep",
			ExtLowerCase: "gitkeep",
			FileName:     ".gitkeep",
			Name:         "",
			Path:         "../../../tests/storages/before/.gitkeep",
		},
		model.File{
			Extension:    "jpg",
			ExtLowerCase: "jpg",
			FileName:     "example1.jpg",
			Name:         "example1",
			Path:         "../../../tests/storages/before/example1.jpg",
		},
		model.File{
			Extension:    "JPG",
			ExtLowerCase: "jpg",
			FileName:     "example2.jpg",
			Name:         "example2",
			Path:         "../../../tests/storages/before/example2.jpg",
		},
	}

	if len(result) != len(expect) {
		t.Error("\nファイル数が違います。")
		t.Error("\n実際: ", len(result), "\n期待: ", len(expect))
		t.Error("\n取得ファイル一覧")
		for _, file := range result {
			t.Error(file.Path + file.Name)
		}
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			if result[i].Extension != expect[i].Extension {
				t.Error("\n実際: ", result[i].Extension, "\n期待: ", expect[i].Extension)
			}
			if result[i].ExtLowerCase != expect[i].ExtLowerCase {
				t.Error("\n実際: ", result[i].ExtLowerCase, "\n期待: ", expect[i].ExtLowerCase)
			}
			if result[i].FileName != expect[i].FileName {
				t.Error("\n実際: ", result[i].FileName, "\n期待: ", expect[i].FileName)
			}
			if result[i].Name != expect[i].Name {
				t.Error("\n実際: ", result[i].Name, "\n期待: ", expect[i].Name)
			}
			if result[i].Path != expect[i].Path {
				t.Error("\n実際: ", result[i].Path, "\n期待: ", expect[i].Path)
			}
		}
	}
}

// TestGetFileListsFromStoragesFailed GetFileListsFromStoragesの異常系テスト
func TestGetFileListsFromStoragesFailed(t *testing.T) {
	fp := di.InjectFilePersistence()
	result, err := fp.GetFileListsFromStorages("../../../tests/storages/nothing")
	if err == nil {
		t.Fatal("failed test")
	}
	if len(result) != 0 {
		t.Fatal("failed test")
	}
}
