package repository

import (
	"testing"

	"github.com/shibayu36/go-playground/diary/config"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	c, _ := config.Load()
	repos, _ := NewRepositories(c.DbDsn)

	err := repos.User.Create("shibayu37@gmail.com", "shibayu37")
	assert.Nil(t, err)
}
