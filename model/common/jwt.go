package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Y1tClaim struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	LoginName  string    `json:"loginName"`
	BufferTime int64     `json:"bufferTime"`
	jwt.StandardClaims
}
