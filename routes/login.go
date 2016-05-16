package routes

import (
	"github.com/tonymtz/gekko/server/status"
	"github.com/tonymtz/gekko/services/oauth2"
	"github.com/tonymtz/gekko/services/oauth2/providers"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var myProviders map[string]oauth2.IProvider

type LoginRoute interface {
	Get(echo.Context) error
	Callback(echo.Context) error
}

var Login LoginRoute

type loginRoute struct {
	LoginRoute
}

func (this *loginRoute) Get(ctx echo.Context) error {
	provider := ctx.Param("provider")

	if p, ok := myProviders[provider]; ok {
		return ctx.Redirect(status.FOUND, p.RedirectUrl())
	}

	return echo.NewHTTPError(status.NOT_FOUND, "Unknown provider")
}

func (this *loginRoute) Callback(ctx echo.Context) error {
	code := ctx.QueryParam("code")

	if code == "" {
		return ctx.String(status.BAD_REQUEST, "Must provide a code")
	}

	provider := ctx.Param("provider")

	if p, ok := myProviders[provider]; ok {
		token, err := p.ExchangeToken(code)

		if err != nil {
			return ctx.String(status.INTERNAL_SERVER_ERROR, "Exchange token error")
		}

		log.Print(token.Token) // TODO - remove this

		// FACT: token exists

		//  gather provider.profile information from provider
		//  if provider.profile data exists
		//      user.getByProviderId(provider.profile.provider_id) in our database
		//      generate custom.token
		//      if models.user exists
		//          update models.user with new provider.token & custom.token
		//      else if models.user doesn't exist
		//          create new models.user with provider.profile & provider.token & custom.token
		//      set cookie with custom.token
		//      redirect to "/app"
		//  else if provider.profile doesn't exist
		//      return error! < weird case
	}

	return ctx.String(200, provider + " " + code)
}

func init() {
	myProviders = make(map[string]oauth2.IProvider)

	myProviders["google"] = providers.NewGoogle(
		"6q4z98mh42d8wqd",
		"rzktu943cdc777r",
		"http://localhost:3000/login/google/callback",
	)

	myProviders["dropbox"] = providers.NewDropbox(
		"6q4z98mh42d8wqd",
		"rzktu943cdc777r",
		"http://localhost:3000/login/dropbox/callback",
	)

	Login = &loginRoute{}
}
