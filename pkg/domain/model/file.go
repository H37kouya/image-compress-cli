package model

// File ファイルクラス
type File struct {
	Dir          string // ディレクトリ
	Extension    string // 拡張子(元)
	ExtLowerCase string // 拡張子(小文字)
	FileName     string // ファイル名(拡張子は小文字)
	Name         string // ファイル名(拡張子なし)
	Path         string // パス
}

// Files ファイルクラス
type Files []File
