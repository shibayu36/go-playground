// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary HTTP client types
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package client

import (
	diary "github.com/shibayu36/go-playground/diary/gen/diary"
)

// UserSignupRequestBody is the type of the "diary" service "UserSignup"
// endpoint HTTP request body.
type UserSignupRequestBody struct {
	// User name
	Name string `form:"name" json:"name" xml:"name"`
	// User email
	Email string `form:"email" json:"email" xml:"email"`
}

// NewUserSignupRequestBody builds the HTTP request body from the payload of
// the "UserSignup" endpoint of the "diary" service.
func NewUserSignupRequestBody(p *diary.UserSignupPayload) *UserSignupRequestBody {
	body := &UserSignupRequestBody{
		Name:  p.Name,
		Email: p.Email,
	}
	return body
}