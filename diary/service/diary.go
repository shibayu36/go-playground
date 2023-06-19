package service

import (
	"context"
	"log"

	diary "github.com/shibayu36/go-playground/diary/gen/diary"
)

// diary service example implementation.
// The example methods log the requests and return zero values.
type diarysrvc struct {
	logger *log.Logger
}

// NewDiary returns the diary service implementation.
func NewDiary(logger *log.Logger) diary.Service {
	return &diarysrvc{logger}
}

// UserSignup implements UserSignup.
func (s *diarysrvc) UserSignup(ctx context.Context, p *diary.UserSignupPayload) (err error) {
	s.logger.Print("diary.UserSignup")
	return
}
