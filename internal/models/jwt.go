package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	User_id int
	Roles   []string
	jwt.StandardClaims
}
