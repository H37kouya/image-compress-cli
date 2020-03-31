/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"image-compress-cli/pkg/di"

	"github.com/spf13/cobra"
)

// resizeImageCmd represents the resizeImage command
var resizeImageCmd = &cobra.Command{
	Use:   "resizeImage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ih := di.InjectImageHandler()

		ih.Index(cmd)
	},
}

func init() {
	// フラグの値をフラグ名で参照する場合
	// 第1引数: フラグ名
	// 第2引数: 短縮フラグ名（末尾が "P" の関数では短縮フラグを指定できる）
	// 第3引数: デフォルト値
	// 第4引数: 説明
	resizeImageCmd.PersistentFlags().IntP("width", "w", 960, "画像の横幅")
	rootCmd.AddCommand(resizeImageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resizeImageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resizeImageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
