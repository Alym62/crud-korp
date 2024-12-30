package jwt

import (
	"fmt"
	"time"

	"github.com/Alym62/crud-korp/internal/models"
	jsonWebToken "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret-key")

func GenerateJWT(user *models.User) (string, error) {
	claims := jsonWebToken.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"position": user.Position,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jsonWebToken.NewWithClaims(jsonWebToken.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseJWT(tokenString string) (*jsonWebToken.Token, error) {
	token, err := jsonWebToken.Parse(tokenString, func(token *jsonWebToken.Token) (interface{}, error) {
		if _, ok := token.Method.(*jsonWebToken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura incorreto: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
