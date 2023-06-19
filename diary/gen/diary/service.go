// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary service
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package diary

import (
	"context"
)

// Service is the diary service interface.
type Service interface {
	// UserSignup implements UserSignup.
	UserSignup(context.Context, *UserSignupPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "diary"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"UserSignup"}

// UserSignupPayload is the payload type of the diary service UserSignup method.
type UserSignupPayload struct {
	// User name
	Name string
	// User email
	Email string
}