package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jstemmer/go-junit-report/parser"
)

func TestJunitFormatter(t *testing.T) {
	filepath := "/Users/shibayu36/Downloads/test-log.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// この時点で、fileはio.Readerインターフェースを実装しています
	reader := io.Reader(file)

	report, _ := parser.Parse(reader, "")
	spew.Dump(report)
}
