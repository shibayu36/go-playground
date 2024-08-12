package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteSize(t *testing.T) {
	// sample.pngから読み込み
	sample, err := os.ReadFile("sample.png")
	require.NoError(t, err)

	// バイトサイズを取得
	byteSize := len(sample)

	// バイトサイズを出力
	fmt.Println(byteSize)
}
