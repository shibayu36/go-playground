package diary

import (
	"context"
	"log"
	"os"
	"testing"

	calc "github.com/shibayu36/go-playground/diary/gen/calc"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	logger := log.New(os.Stderr, "LOG: ", log.Lshortfile)
	service := NewCalc(logger)
	ctx := context.Background()

	tests := []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{
			name:   "Add positive numbers",
			inputA: 5,
			inputB: 3,
			want:   8,
		},
		{
			name:   "Add negative numbers",
			inputA: -2,
			inputB: -3,
			want:   -5,
		},
		{
			name:   "Add zero",
			inputA: 0,
			inputB: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Add(ctx, &calc.AddPayload{A: tt.inputA, B: tt.inputB})
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
