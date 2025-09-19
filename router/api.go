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
	api.Echo.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	api.Echo.POST("/admin/login", api.UserHandler.HandlerLogin)
	api.Echo.POST("/admin/signup", api.UserHandler.HandlerSignUp)

	// Route yêu cầu xác thực JWT
	api.Echo.GET("/admin/profile", api.UserHandler.Profile, middleware.JWTMiddleware())

	// Route cho users - chỉ admin mới truy cập được
	api.Echo.GET("/admin/users", api.UserHandler.GetAllUsers, middleware.JWTMiddleware())

	// Route cho infra components - chỉ admin mới truy cập được
	api.Echo.GET("/admin/infra-components", api.InfraComponentHandler.GetInfraComponents, middleware.JWTMiddleware())
	api.Echo.GET("/admin/infra-components/all", api.InfraComponentHandler.GetAllInfraComponents, middleware.JWTMiddleware())
	api.Echo.GET("/admin/infra-components/pending", api.InfraComponentHandler.GetPendingInfraComponents, middleware.JWTMiddleware())

	// Route cập nhật status infra components - chỉ admin mới được sửa
	api.Echo.PUT("/admin/infra-components/status", api.InfraComponentHandler.UpdateInfraComponentStatus, middleware.JWTMiddleware())

	// Route cập nhật thông tin infra components - chỉ admin mới được sửa
	api.Echo.PUT("/admin/infra-components", api.InfraComponentHandler.UpdateInfraComponent, middleware.JWTMiddleware())
}
