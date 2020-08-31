package jwt

import (
	"gin-example/pkg/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.JwtConfig.Secret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Rolename string `json:"rolename"`
	jwt.StandardClaims
}

func GenerateToken(username, password, rolename string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(10 * time.Second)

	claims := Claims{
		username,
		password,
		rolename,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
