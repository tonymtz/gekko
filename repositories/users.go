package repositories

import (
	"database/sql"
	"github.com/tonymtz/gekko/models"
	"errors"
	"log"
)

type UsersRepository interface {
	FindById(int) (*models.User, error)
	FindByProviderId(string) (*models.User, error)
	Insert(*models.User) (int, error)
	Remove()
	Update(*models.User)
}

type userRepository struct {
	database *sql.DB
}

func NewUsersRepository(database *sql.DB) UsersRepository {
	return &userRepository{
		database: database,
	}
}

func (this *userRepository) Insert(newUser *models.User) (int, error) {
	var lastInsertedId int

	err := this.database.QueryRow(
		"INSERT INTO users (id_provider, display_name, email, profile_picture, role, token, jwt) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		newUser.IdProvider,
		newUser.DisplayName,
		newUser.Email,
		newUser.ProfilePicture,
		newUser.Role,
		newUser.Token,
		newUser.JWT,
	).Scan(&lastInsertedId)

	if err != nil {
		return -1, err
	}

	return lastInsertedId, nil
}

func (this *userRepository) FindById(id int) (*models.User, error) {
	foundUser := &models.User{}

	err := this.database.QueryRow(
		"SELECT id, id_provider, display_name, email, profile_picture, role, token, jwt FROM users WHERE id=$1",
		id,
	).Scan(
		&foundUser.Id,
		&foundUser.IdProvider,
		&foundUser.DisplayName,
		&foundUser.Email,
		&foundUser.ProfilePicture,
		&foundUser.Role,
		&foundUser.Token,
		&foundUser.JWT,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("No user with that ID.")
	}

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (this *userRepository) FindByProviderId(providerId string) (*models.User, error) {
	foundUser := &models.User{}

	err := this.database.QueryRow(
		"SELECT id, id_provider, display_name, email, profile_picture, role, token, jwt FROM users WHERE id_provider=$1",
		providerId,
	).Scan(
		&foundUser.Id,
		&foundUser.IdProvider,
		&foundUser.DisplayName,
		&foundUser.Email,
		&foundUser.ProfilePicture,
		&foundUser.Role,
		&foundUser.Token,
		&foundUser.JWT,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("No user with that ID.")
	}

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (this *userRepository) Remove() {
	log.Panic("Not implemented")
}

func (this *userRepository) Update(*models.User) {
	log.Panic("Not implemented")
}
