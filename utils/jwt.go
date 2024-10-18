package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "os"
    "time"
)

func GenerateJWT(userID uint) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
    secret := os.Getenv("JWT_SECRET")
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    return token, err
}
