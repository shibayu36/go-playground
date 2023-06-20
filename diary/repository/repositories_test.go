package repository

import (
	"testing"

	"github.com/shibayu36/go-playground/diary/config"
	"github.com/stretchr/testify/assert"
)

func TestNewRepositories(t *testing.T) {
	c, _ := config.Load()
	repos, err := NewRepositories(c.DbDsn)

	assert.Nil(t, err)
	assert.Nil(t, repos.db.Ping(), "db should be connected")
}

func TestClose(t *testing.T) {
	c, _ := config.Load()
	repos, _ := NewRepositories(c.DbDsn)

	assert.Nil(t, repos.Close())
	assert.NotNil(t, repos.db.Ping(), "db should be disconnected")
}
