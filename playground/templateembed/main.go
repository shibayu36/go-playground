package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
)

//go:embed sample.json.tmpl
var sampleJsonTmpl string

func main() {
	fmt.Println(sampleJsonTmpl)

	tmpl, err := template.New("sample").Parse(sampleJsonTmpl)
	if err != nil {
		os.Exit(1)
	}

	data := map[string]any{
		"Hoge": "hogehoge",
		"Foo":  "barbar",
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		os.Exit(1)
	}
	fmt.Println(buf.String())
}
