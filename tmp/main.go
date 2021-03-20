/*
@Time : 2021/3/20 13:32
@Author : Tux
@Description :
*/

package main

import (
	"html/template"
	"os"
	"strings"
)

const templateText = `
Output 0: {{ title .Name1 }}
Output 0: {{ title .Name2 }}
Output 0: {{ .Name3 |title }}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data:= map[string]string{
		"Name1": "go",
		"Name2":"programing",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
