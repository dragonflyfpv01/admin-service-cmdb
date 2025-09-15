package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sllpklls/template-backend-go/model"
)

const SECRET_KEY = "hoangthaifc01"

// func JWTMiddleware
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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
		return SECRET_KEY, nil
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
