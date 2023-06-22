package service

import (
	"context"
	"log"

	diary "github.com/shibayu36/go-playground/diary/gen/diary"
	"github.com/shibayu36/go-playground/diary/repository"
)

// diary service example implementation.
// The example methods log the requests and return zero values.
type diarysrvc struct {
	logger *log.Logger
	repos  *repository.Repositories
}

// NewDiary returns the diary service implementation.
func NewDiary(logger *log.Logger, repos *repository.Repositories) diary.Service {
	return &diarysrvc{logger, repos}
}

// UserSignup implements UserSignup.
func (s *diarysrvc) UserSignup(ctx context.Context, p *diary.UserSignupPayload) (err error) {
	// TODO:
	// * Email Length, Format
	// * Name Length, Name Available Characters
	// * Duplication of Email or Name
	_, err = s.repos.User.Create(p.Email, p.Name)
	return
}
