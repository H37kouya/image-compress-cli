package model

// File ファイルクラス
type File struct {
	Extension    string // 拡張子(元)
	ExtLowerCase string // 拡張子(小文字)
	FileName     string // ファイル名
	Name         string // ファイル名(拡張子なし)
	Path         string // パス
}

// Files ファイルクラス
type Files []File
