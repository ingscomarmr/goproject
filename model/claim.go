package model

import jwt "github.com/dgrijalva/jwt-go"

//modelo auxiliar para validar si se autentico por token
type Claim struct {
	User User `json:"user"`
	jwt.StandardClaims
}
