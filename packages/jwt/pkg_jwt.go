package pkgjwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webtoken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return webtoken, err
}

// function fo verification if user login with JWT
func VerifyToken(tokenString string) (*jwt.Token, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing methode %v", t.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, err
}

// function for decode similarity JWT if user has been login
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
