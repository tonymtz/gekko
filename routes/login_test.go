package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/tonymtz/gekko/repos"
	"github.com/tonymtz/gekko/services"
	"github.com/tonymtz/gekko/test_utils/mocks"
)

var (
	sampleProfile = &services.GoogleProfile{
		Id: "gid:111",
		DisplayName: "test user",
		Emails: []*services.GoogleProfileEmail{
			&services.GoogleProfileEmail{
				Value: "test@sample.com",
				Type: "account",
			},
		},
		Image: &services.GoogleProfileImage{
			Url: "http://sample.com/image.png",
		},
	}
)

func TestLoginRoute_Get(t *testing.T) {
	// Setup
	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	myLoginRoute := &loginRoute{}

	// Negative Case
	if err := myLoginRoute.Get(ctx); err == nil {
		t.Error(err)
	}

	// Positive Case
	ctx = e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	ctx.SetParamNames("provider")
	ctx.SetParamValues("google")
	if err := myLoginRoute.Get(ctx); err != nil {
		t.Error(err)
	}

	// Positive Case
	ctx = e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	ctx.SetParamNames("provider")
	ctx.SetParamValues("dropbox")
	if err := myLoginRoute.Get(ctx); err != nil {
		t.Error(err)
	}
}

func TestLoginRoute_Callback(t *testing.T) {
	// Setup
	e := echo.New()
	req := new(http.Request)

	newUrl, _ := url.Parse("?code=my_code")
	req.URL = newUrl

	rec := httptest.NewRecorder()
	ctx := e.NewContext(
		standard.NewRequest(req, e.Logger()),
		standard.NewResponse(rec, e.Logger()),
	)

	ctx.SetParamNames("provider")
	ctx.SetParamValues("google")

	var userRepoFake repos.UsersRepository
	var myLoginRoute LoginRoute
	var googleAPIMock services.GoogleAPI

	// Positive Case
	userRepoFake = mocks.NewUsersRepoFake(true)
	googleAPIMock = mocks.NewGoogleAPIMock(false, sampleProfile)

	myLoginRoute = &loginRoute{
		usersRepository: userRepoFake,
		googleAPI: googleAPIMock,
	}

	err := myLoginRoute.Callback(ctx)

	if err != nil {
		t.Error(err)
	}

	// Negative Case
	ctx = e.NewContext(
		standard.NewRequest(req, e.Logger()),
		standard.NewResponse(rec, e.Logger()),
	)

	ctx.SetParamNames("provider")
	ctx.SetParamValues("otherNotListedProvider")

	userRepoFake = mocks.NewUsersRepoFake(true)
	googleAPIMock = mocks.NewGoogleAPIMock(false, sampleProfile)

	myLoginRoute = &loginRoute{
		usersRepository: userRepoFake,
		googleAPI: googleAPIMock,
	}

	err = myLoginRoute.Callback(ctx)

	if err == nil {
		t.Error(err)
	}
}
