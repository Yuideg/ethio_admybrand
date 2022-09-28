package model

import (
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Claim struct {
	UserID    uuid.UUID `json:"user_id"`
	Authorize bool      `json:"authorized"`
	Username  string    `json:"username"`
	jwt.StandardClaims
}
