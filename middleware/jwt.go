package middleware

import (
	"sllpklls/admin-service/model"
	"sllpklls/admin-service/security"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware tạo middleware để xác thực JWT
func JWTMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(security.SECRET_KEY), // Sử dụng secret key từ security package
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &model.JwtCustomClaims{}
		},
	}
	return echojwt.WithConfig(config)
}
