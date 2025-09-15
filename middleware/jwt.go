package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sllpklls/template-backend-go/model"
	"github.com/sllpklls/template-backend-go/security"
)

// JWTMiddleware tạo middleware để xác thực JWT
func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(security.SECRET_KEY), // Sử dụng secret key từ security package
		Claims:     &model.JwtCustomClaims{},    // Sử dụng struct JwtCustomClaims từ model
	}
	return middleware.JWTWithConfig(config)
}
