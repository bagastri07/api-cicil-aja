package token

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID       uint64 `json:"ID"`
	Email    string `json:"email"`
	UserType string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id uint64, email string, userType string) (string, error) {

	claims := &JwtCustomClaims{
		ID:       id,
		Email:    email,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return t, nil
}

func DecodeToken(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	return claims
}
