package security

import (
	"errors"
	"time"

	"sllpklls/admin-service/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const SECRET_KEY = "hoangthaifc01"

// func JWTMiddleware
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
func validateToken(tokenString string) (*model.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JwtCustomClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

// GetClaimsFromContext lấy JWT claims từ context của Echo
func GetClaimsFromContext(c echo.Context) (*model.JwtCustomClaims, error) {
	user := c.Get("user")
	if user == nil {
		return nil, errors.New("user not found in context")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return nil, errors.New("invalid token format")
	}

	claims, ok := token.Claims.(*model.JwtCustomClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	return claims, nil
}

// IsAdmin kiểm tra xem user có role admin hay không
func IsAdmin(claims *model.JwtCustomClaims) bool {
	return claims.Role == model.ADMIN.String()
}
