package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithType(t *testing.T) {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
