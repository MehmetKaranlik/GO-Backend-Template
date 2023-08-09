package jwt

import (
	"Backend/Core/Globals"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	access_token_duration  = time.Hour * 24
	refresh_token_duration = time.Hour * 24 * 30
	verify_token_duration  = time.Minute * 30
	key_exp                = "exp"
	key_id                 = "id"
)

func GenerateAccessToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key_exp: time.Now().Add(access_token_duration).Unix(),
		key_id:  id,
	})

	// TODO: AccessSecretKey should be in a config file
	tokenString, err := token.SignedString([]byte(Globals.EnvValues.JWTSecrets.AccessSecretKey))
	if err != nil {
		fmt.Printf("\nError while generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func GenerateRefreshToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key_exp: time.Now().Add(refresh_token_duration).Unix(),
		key_id:  id,
	})
	tokenString, err := token.SignedString([]byte(Globals.EnvValues.JWTSecrets.RefreshSecretKey))
	if err != nil {
		fmt.Printf("\nError while generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func ParseAccessToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Globals.EnvValues.JWTSecrets.AccessSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}

	return "", nil
}

func GenerateVerifyToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		key_exp: time.Now().Add(verify_token_duration).Unix(),
		key_id:  id,
	})

	tokenString, err := token.SignedString([]byte(Globals.EnvValues.JWTSecrets.VerifySecretKey))
	if err != nil {
		fmt.Printf("\nError while generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ParseVerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Globals.EnvValues.JWTSecrets.VerifySecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}

	return "", nil
}
