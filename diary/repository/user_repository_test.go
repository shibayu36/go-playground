package repository

import (
	"testing"

	"github.com/shibayu36/go-playground/diary/config"
	"github.com/shibayu36/go-playground/diary/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	c, _ := config.Load()
	repos, _ := NewRepositories(c.DbDsn)

	email := testutil.RandomEmail()
	name := testutil.RandomString(10)

	user, err := repos.User.Create(
		email, name,
	)
	assert.Nil(t, err)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, name, user.Name)
}

func TestUserRepositoryFindByID(t *testing.T) {
	c, _ := config.Load()
	repos, _ := NewRepositories(c.DbDsn)

	email := testutil.RandomEmail()
	name := testutil.RandomString(10)
	user, _ := repos.User.Create(
		email, name,
	)

	foundUser, err := repos.User.FindByID(user.UserID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Name, foundUser.Name)
}
