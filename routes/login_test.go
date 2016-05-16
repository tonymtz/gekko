package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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
