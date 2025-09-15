package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sllpklls/template-backend-go/handler"
	"github.com/sllpklls/template-backend-go/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	// Route không yêu cầu xác thực JWT
	api.Echo.POST("/user/sign-in", api.UserHandler.HandlerSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
