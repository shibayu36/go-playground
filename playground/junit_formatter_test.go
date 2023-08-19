package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jstemmer/go-junit-report/v2/junit"
	parser "github.com/jstemmer/go-junit-report/v2/parser/gotest"
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

	report, _ := parser.NewParser().Parse(reader)
	// spew.Dump(report)
	testsuites := junit.CreateFromReport(report, "")

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	if err := enc.Encode(testsuites); err != nil {

	}
	if err := enc.Flush(); err != nil {

	}
	fmt.Fprintf(os.Stdout, "\n")
}
