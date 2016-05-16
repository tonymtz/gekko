package mocks

import (
	"github.com/tonymtz/gekko/services"
	"errors"
)

type GoogleAPIMock struct {
	services.GoogleAPI
	throwError bool
	profile    *services.GoogleProfile
}

func NewGoogleAPIMock(throwError bool, profile *services.GoogleProfile) services.GoogleAPI {
	return &GoogleAPIMock{
		throwError: throwError,
		profile: profile,
	}
}

func (this *GoogleAPIMock) GetProfile(token string) (*services.GoogleProfile, error) {
	if this.throwError {
		return nil, errors.New("Something went wrong")
	}

	return this.profile, nil
}
