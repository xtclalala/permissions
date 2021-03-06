package system

import (
	"github.com/dgrijalva/jwt-go"
	"permissions/global"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
	"time"
)

type JwtService struct {
	SignKey []byte
}

func NewJWT() *JwtService {
	return &JwtService{
		[]byte(global.System.Jwt.SignKey),
	}
}

// CreateClaim 创建claim
func (j *JwtService) CreateClaim(user *system.SysUser) common.Y1tClaim {
	jwtConfig := global.System.Jwt
	timeout := jwtConfig.Timeout
	bufferTime := jwtConfig.BufferTime
	iss := jwtConfig.Iss
	return common.Y1tClaim{
		Id:         user.ID,
		Username:   user.Username,
		LoginName:  user.LoginName,
		BufferTime: bufferTime * 60 * 60,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + timeout*60*60,
			Issuer:    iss,
		},
	}
}

// CreateJwt 创建jwt
func (j *JwtService) CreateJwt(clain *common.Y1tClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)
	return token.SignedString(j.SignKey)
}

// CreateJwtByOldClaim 生成新的jwt
func (j *JwtService) CreateJwtByOldClaim(claim *common.Y1tClaim) (string, error) {
	return j.CreateJwt(claim)
}

// ParseJwt 解析jwt
func (j *JwtService) ParseJwt(tokenString string) (*common.Y1tClaim, int) {
	token, err := jwt.ParseWithClaims(tokenString, &common.Y1tClaim{}, func(token *jwt.Token) (i any, err error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, utils.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, utils.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, utils.TokenNotValid
			} else {
				return nil, utils.TokenInvalid
			}
		}
	}
	if token != nil {
		if claim, ok := token.Claims.(*common.Y1tClaim); ok && token.Valid {
			return claim, 0
		} else {
			return nil, utils.TokenInvalid
		}
	}
	return nil, utils.TokenInvalid
}
