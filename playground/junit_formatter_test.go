package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jstemmer/go-junit-report/formatter"
	parser "github.com/jstemmer/go-junit-report/parser"
	"github.com/jstemmer/go-junit-report/v2/junit"
	parserv2 "github.com/jstemmer/go-junit-report/v2/parser/gotest"
	"github.com/stretchr/testify/require"
)

func TestJunitFormatter(t *testing.T) {
	filepath := "/Users/shibayu36/Downloads/test-log-minimum.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	reader := io.Reader(file)

	report, _ := parser.Parse(reader, "")

	err = formatter.JUnitReportXML(report, false, "", os.Stdout)
	require.NoError(t, err)
}

func TestJunitFormatterV2(t *testing.T) {
	filepath := "/Users/shibayu36/Downloads/test-log-minimum.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	reader := io.Reader(file)

	report, _ := parserv2.NewParser().Parse(reader)
	// spew.Dump(report)
	testsuites := junit.CreateFromReport(report, "")

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "\t")
	if err := enc.Encode(testsuites); err != nil {
		t.Fatal(err)
	}
	if err := enc.Flush(); err != nil {
		t.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "\n")
}
