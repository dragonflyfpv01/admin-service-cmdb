package router

import (
	"sllpklls/admin-service/handler"
	"sllpklls/admin-service/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo                  *echo.Echo
	UserHandler           handler.UserHandler
	InfraComponentHandler handler.InfraComponentHandler
}

func (api *API) SetupRouter() {
	// Route không yêu cầu xác thực JWT
	api.Echo.POST("/admin/login", api.UserHandler.HandlerLogin)
	api.Echo.POST("/admin/signup", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/admin/profile", api.UserHandler.Profile, middleware.JWTMiddleware())

	// Route cho infra components - chỉ admin mới truy cập được
	api.Echo.GET("/admin/infra-components", api.InfraComponentHandler.GetInfraComponents, middleware.JWTMiddleware())
}
