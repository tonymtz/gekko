package mocks

import (
	"github.com/tonymtz/gekko/repos"
	"github.com/tonymtz/gekko/models"
	"errors"
)

type UsersRepoFake struct {
	repos.UsersRepository
	throwError bool
}

func NewUsersRepoFake(throwError bool) repos.UsersRepository {
	return &UsersRepoFake{
		throwError: throwError,
	}
}

func (this *UsersRepoFake) FindById(id int64) (*models.User, error) {
	if this.throwError {
		return nil, errors.New("Something went wrong")
	}

	user := &models.User{}

	return user, nil
}

func (this *UsersRepoFake) FindByProviderId(providerId string) (*models.User, error) {
	if this.throwError {
		return nil, errors.New("Something went wrong")
	}

	user := &models.User{}

	return user, nil
}

func (this *UsersRepoFake) Insert(newUser *models.User) (int64, error) {
	if this.throwError {
		return 0, errors.New("Something went wrong")
	}

	return 0, nil
}

func (this *UsersRepoFake) Remove() {}

func (this *UsersRepoFake) Update(newUser *models.User) {}
