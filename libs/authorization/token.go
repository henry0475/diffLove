package authorization

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func createToken(uid string, expiredHour int) (tokenString string, err error) {
	// provide a token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": uid,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(expiredHour)).Unix(),
		"sub": "AuthToken",
	})
	return token.SignedString(privateKey)
}

func getUIDFromToken(tokenString string) (uid string, err error) {
	// validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["uid"].(string), nil
	}

	return "", errors.New("Error: It is not a valid token")
}

func refreshToken(oldToken string) (tokenString string, err error) {
	// validate the token
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return createToken(claims["uid"].(string), conf.Validation.Token.App)
	}

	return "", errors.New("ERROR: Token is wrong")
}
