package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	require.NoError(t, err)
	assert.Equal(t, wantPort, got.Port)
	assert.Equal(t, "dev", got.Env)
}
