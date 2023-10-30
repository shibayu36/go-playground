package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	err := json.Unmarshal(want, &jw)
	require.NoError(t, err)
	err = json.Unmarshal(got, &jg)
	require.NoError(t, err)

	if diff := cmp.Diff(jg, jw); diff != "" {
		t.Errorf("json mismatch (-want +got):\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })

	gb, err := io.ReadAll(got.Body)
	require.NoError(t, err)

	assert.Equal(t, status, got.StatusCode)

	if len(gb) == 0 && len(body) == 0 {
		return
	}

	AssertJSON(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	require.NoError(t, err)
	return bt
}
