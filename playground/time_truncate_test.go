package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTruncate(t *testing.T) {
	now := time.Date(2025, 2, 18, 12, 0, 3, 0, time.Local)
	truncated1 := now.Truncate(time.Minute * 10)
	fmt.Println(truncated1)
	truncated2 := now.Add(time.Minute * 5).Truncate(time.Minute * 10)
	fmt.Println(truncated2)
	truncated3 := now.Add(time.Minute * 15).Truncate(time.Minute * 10)
	fmt.Println(truncated3)
}
