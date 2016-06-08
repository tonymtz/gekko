package routes

import (
	"github.com/tonymtz/gekko/models"
	"github.com/tonymtz/gekko/repos"
	"github.com/tonymtz/gekko/server/config"
	"github.com/tonymtz/gekko/server/status"
	"github.com/tonymtz/gekko/services"
	"github.com/tonymtz/gekko/services/oauth2"
	"github.com/tonymtz/gekko/services/oauth2/providers"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

var myProviders map[string]oauth2.IProvider

type LoginRoute interface {
	Get(echo.Context) error
	Callback(echo.Context) error
}

var Login LoginRoute

type loginRoute struct {
	LoginRoute
	usersRepository repos.UsersRepository
	googleAPI       services.GoogleAPI
}

func (this *loginRoute) Get(ctx echo.Context) error {
	provider := ctx.Param("provider")

	if p, ok := myProviders[provider]; ok {
		return ctx.Redirect(status.FOUND, p.RedirectUrl())
	}

	return echo.NewHTTPError(status.NOT_FOUND, "Unknown provider")
}

func (this *loginRoute) Callback(ctx echo.Context) error {
	provider := ctx.Param("provider")

	var myProvider = myProviders[provider];

	code := ctx.QueryParam("code")

	if code == "" {
		return echo.NewHTTPError(status.BAD_REQUEST, "Must provide a code")
	}

	tokenFromProvider, err := myProvider.ExchangeToken(code)

	if err != nil {
		return echo.NewHTTPError(status.INTERNAL_SERVER_ERROR, "Exchange token error")
	}

	myGoogleProfile := this.retrieveProfileFromTokenAndUpdateUser(tokenFromProvider.Token)

	mySessionProfile := &SessionProfile{
		Email: myGoogleProfile.Emails[0].Value,
		Token: tokenFromProvider.Token,
	}

	myJWT := this.createJWTFromProfile(mySessionProfile)

	return ctx.JSON(status.OK, &CallbackResponse{
		Token: myJWT,
	})
}

func (this *loginRoute) retrieveProfileFromTokenAndUpdateUser(token string) *services.GoogleProfile {
	profile, _ := this.googleAPI.GetProfile(token)
	user, err := this.usersRepository.FindByProviderId(profile.Id)

	if err == nil {
		// update
		user.Token = token
		this.usersRepository.Update(user)
	} else {
		// create
		this.usersRepository.Insert(
			&models.User{
				IdProvider: profile.Id,
				Email: profile.Emails[0].Value,
				DisplayName: profile.DisplayName,
				Token: token,
			},
		)
	}

	return profile
}

func (this *loginRoute) createJWTFromProfile(profile *SessionProfile) string {
	jwtoken := jwt.New(jwt.SigningMethodHS256)

	jwtoken.Claims["email"] = profile.Email
	jwtoken.Claims["token"] = profile.Token

	t, _ := jwtoken.SignedString([]byte("secret"))

	return t
}

func init() {
	myProviders = make(map[string]oauth2.IProvider)

	myProviders["google"] = providers.NewGoogle(
		config.Config.GoogleId,
		config.Config.GoogleSecret,
		config.Config.GoogleCallback,
	)

	myProviders["dropbox"] = providers.NewDropbox(
		config.Config.DropboxId,
		config.Config.DropboxSecret,
		config.Config.DropboxCallback,
	)

	Login = &loginRoute{
		usersRepository: repos.NewUsersRepository(config.Config.Database),
		googleAPI: services.NewGoogleAPI(),
	}
}

type SessionProfile struct {
	Email string
	Token string
}

type CallbackResponse struct {
	Token string `json:"token"`
}
