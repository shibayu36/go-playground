package model

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	type args struct {
		email string
		name  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"valid",
			args{
				email: "shibayu36@gmail.com",
				name:  "shibayu36",
			},
			nil,
		},
		{
			"email is too long",
			args{
				email: strings.Repeat("a", 256) + "@gmail.com",
				name:  "shibayu36",
			},
			errors.New("email is too long"),
		},
		{
			"email is invalid",
			args{
				email: "shibayu36",
				name:  "shibayu36",
			},
			errors.New("email is invalid"),
		},
		{
			"name is too short",
			args{
				email: "shibayu36@gmail.com",
				name:  "sh",
			},
			errors.New("name is too short"),
		},
		{
			"name is too long",
			args{
				email: "shibayu36@gmail.com",
				name:  strings.Repeat("a", 256),
			},
			errors.New("name is too long"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUser(tt.args.email, tt.args.name)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
