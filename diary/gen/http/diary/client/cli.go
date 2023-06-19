// Code generated by goa v3.11.3, DO NOT EDIT.
//
// diary HTTP client CLI support package
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package client

import (
	"encoding/json"
	"fmt"

	diary "github.com/shibayu36/go-playground/diary/gen/diary"
)

// BuildUserSignupPayload builds the payload for the diary UserSignup endpoint
// from CLI flags.
func BuildUserSignupPayload(diaryUserSignupBody string) (*diary.UserSignupPayload, error) {
	var err error
	var body UserSignupRequestBody
	{
		err = json.Unmarshal([]byte(diaryUserSignupBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Laudantium repellendus.\",\n      \"name\": \"Aliquid doloribus.\"\n   }'")
		}
	}
	v := &diary.UserSignupPayload{
		Name:  body.Name,
		Email: body.Email,
	}

	return v, nil
}