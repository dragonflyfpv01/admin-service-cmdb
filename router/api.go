package router

import (
	"sllpklls/admin-service/handler"
	"sllpklls/admin-service/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	// Route không yêu cầu xác thực JWT
	api.Echo.POST("/admin/login", api.UserHandler.HandlerLogin)
	api.Echo.POST("/admin/signup", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/admin/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
