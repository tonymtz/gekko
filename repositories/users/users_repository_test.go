package users

import (
	"testing"
	"github.com/tonymtz/gekko/models"
	"github.com/stretchr/testify/assert"
	"database/sql"
	"errors"
	"log"
	_ "github.com/lib/pq"
)

var (
	testUser = &models.User{
		IdProvider: "uid:111111",
		DisplayName: "test user",
		Email: "test@sample.com",
		ProfilePicture: "http://image.url/",
		Role: 1,
		Token: "t11111",
		JWT: "xxx",
	}

	demoUser = &models.User{
		IdProvider: "uid:222222",
		DisplayName: "demo user",
		Email: "demo@sample.com",
		ProfilePicture: "http://image.net/",
		Role: 1,
		Token: "t22222",
		JWT: "yyy",
	}
)

func getTestDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=welcome1 database=gekko_db")

	if err != nil {
		return nil, err
	}

	db.Exec("DELETE FROM users")

	return db, nil
}

func TestUserRepository_Insert(t *testing.T) {
	/*** SETUP ***/

	db, err := getTestDB()

	if err != nil {
		log.Panic(err.Error())
	}

	userRepo := NewUsersRepository(db)
	userFromDB := &models.User{}

	/*** TEST ***/

	lastId, err := userRepo.Insert(testUser)

	if assert.Nil(t, err) {
		assert.True(t, lastId > -1, "returned id should be greater than -1")

		getUser(db, lastId, userFromDB)

		assert.EqualValues(t, userFromDB, testUser, "both users should be equals")
	}
}

func TestUserRepository_FindById(t *testing.T) {
	/*** SETUP ***/

	db, err := getTestDB()

	if err != nil {
		log.Panic(err.Error())
	}

	first_id, _ := insertUser(db, testUser)
	second_id, _ := insertUser(db, demoUser)

	userRepo := NewUsersRepository(db)

	/*** TEST ***/

	userFromRepo, err := userRepo.FindById(first_id)

	userFromRepo.Id = 0 // our fixture does not have id

	if assert.Nil(t, err) {
		assert.EqualValues(t, userFromRepo, testUser, "both users should be equals")
	}

	userFromRepo, err = userRepo.FindById(second_id)

	userFromRepo.Id = 0 // our fixture does not have id

	if assert.Nil(t, err) {
		assert.EqualValues(t, userFromRepo, demoUser, "both users should be equals")
	}
}

func TestUserRepository_FindByProviderId(t *testing.T) {
	/*** SETUP ***/

	db, err := getTestDB()

	if err != nil {
		log.Panic(err.Error())
	}

	insertUser(db, testUser)
	insertUser(db, demoUser)

	userRepo := NewUsersRepository(db)

	/*** TEST ***/

	userFromRepo, err := userRepo.FindByProviderId("uid:111111")

	userFromRepo.Id = 0 // our fixture does not have id

	if assert.Nil(t, err) {
		assert.EqualValues(t, userFromRepo, testUser, "both users should be equals")
	}

	userFromRepo, err = userRepo.FindByProviderId("uid:222222")

	userFromRepo.Id = 0 // our fixture does not have id

	if assert.Nil(t, err) {
		assert.EqualValues(t, userFromRepo, demoUser, "both users should be equals")
	}
}

func getUser(database *sql.DB, id int, user *models.User) error {
	err := database.QueryRow(
		"SELECT id_provider, display_name, email, profile_picture, role, token, jwt FROM users WHERE id=$1",
		id,
	).Scan(
		&user.IdProvider,
		&user.DisplayName,
		&user.Email,
		&user.ProfilePicture,
		&user.Role,
		&user.Token,
		&user.JWT,
	)

	if err == sql.ErrNoRows {
		return errors.New("No user with that ID.")
	}

	if err != nil {
		return err
	}

	return nil
}

func insertUser(database *sql.DB, user *models.User) (int, error) {
	var lastInsertedId int

	err := database.QueryRow(
		"INSERT INTO users (id_provider, display_name, email, profile_picture, role, token, jwt) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.IdProvider,
		user.DisplayName,
		user.Email,
		user.ProfilePicture,
		user.Role,
		user.Token,
		user.JWT,
	).Scan(&lastInsertedId)

	if err != nil {
		return -1, err
	}

	return lastInsertedId, nil
}
