package repository

import (
	"testing"

	"github.com/Songmu/flextime"
	"github.com/shibayu36/go-playground/diary/config"
	"github.com/shibayu36/go-playground/diary/model"
	"github.com/shibayu36/go-playground/diary/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	t.Run("with valid parameters", func(t *testing.T) {
		restore := flextime.Fix(flextime.Now())
		defer restore()

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
		assert.Equal(t, flextime.Now(), user.CreatedAt)
		assert.Equal(t, flextime.Now(), user.UpdatedAt)

		foundUser, _ := repos.User.FindByID(user.UserID)
		assert.Equal(t, user.UserID, foundUser.UserID, "user is created correctly")
	})

	t.Run("with invalid parameters", func(t *testing.T) {
		restore := flextime.Fix(flextime.Now())
		defer restore()

		c, _ := config.Load()
		repos, _ := NewRepositories(c.DbDsn)

		email := "invalidemail"
		name := testutil.RandomString(10)

		_, err := repos.User.Create(
			email, name,
		)

		assert.Equal(t, model.NewValidationError("email is invalid"), err)
	})
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
