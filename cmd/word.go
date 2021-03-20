/*
@Time : 2021/3/20 10:49
@Author : Tux
@Description :
*/

package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"chapter01/internal/word"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelcase
	ModeUnderscoreToLowerCamelcase
	ModeCamelcaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式的转换,模式如下: ",
	"1: 全部单词转为大写",
	"2: 全部单词转换为小写",
	"3: 下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"5: 驼峰单词转换为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatal("暂不支持改转换格式, 请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s\n", content)
	},
}

var str string
var mode int8
func init() {
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入单词转换的格式")
}
