package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonizeTimestamp(t *testing.T) {
	jsonStr := `{
		"id": 1,
		"name": "test",
		"created_at": "2020-01-01T00:00:00Z",
		"updated_at": "2020-01-01T00:00:00Z"
	}`

	type commonizeTimestamp struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type commonizeStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		commonizeTimestamp
	}

	var hoge commonizeStruct

	fmt.Println(reflect.ValueOf(hoge))
	st := reflect.TypeOf(hoge)
	for i := 0; i < st.NumField(); i++ {
		fmt.Println(st.Field(i).Name)
	}

	err := json.Unmarshal([]byte(jsonStr), &hoge)
	require.NoError(t, err)

	// 構造体がネストされていたとしても、
	assert.Equal(t, 1, hoge.ID)
	assert.Equal(t, "test", hoge.Name)
	assert.Equal(t, "2020-01-01T00:00:00Z", hoge.CreatedAt)
	assert.Equal(t, "2020-01-01T00:00:00Z", hoge.UpdatedAt)
}
