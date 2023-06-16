package diary

import (
	"context"
	"log"

	user "github.com/shibayu36/go-playground/diary/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Signup implements signup.
func (s *usersrvc) Signup(ctx context.Context, p *user.SignupPayload) (err error) {
	s.logger.Print("user.signup")
	return
}
