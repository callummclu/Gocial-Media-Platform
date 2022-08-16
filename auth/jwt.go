package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = username
	claims["aud"] = "go-social.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		err = errors.New("something went wrong")
		return "", err
	}
	return newToken, nil
}

func CheckJWT(token string, username *string) error {
	t, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return err
	} else if !t.Valid {
		err = errors.New("Invalid token")
		return err
	} else if t.Claims.(jwt.MapClaims)["aud"] != "go-social.jwtgo.io" {
		err = errors.New("Invalid aud")
		return err
	} else if t.Claims.(jwt.MapClaims)["iss"] != "jwtgo.io" {
		err = errors.New("Invalid iss")
		return err
	}

	claims := t.Claims.(jwt.MapClaims)
	*username = claims["user"].(string)

	return nil
}
