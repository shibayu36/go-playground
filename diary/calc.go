package calcapi

import (
	"context"
	"log"

	calc "github.com/shibayu36/go-playground/diary/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return
}