package middleware

import (
	"github.com/gin-gonic/gin"
	"permissions/global"
	"permissions/model/common"
	"permissions/services/system"
	"permissions/utils"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := global.System.App.Auth
		token := c.Request.Header.Get(auth)
		if token == "" {
			common.FailWhitStatus(utils.NOTOKEN, c)
			c.Abort()
			return
		}

		j := system.NewJWT()
		claims, err := j.ParseJwt(token)
		if err != 0 {
			common.FailWhitStatus(err, c)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.System.Jwt.Timeout
			t, _ := j.CreateJwt(claims)
			c.Header(auth, t)
		}
		c.Set("claims", claims)
		c.Next()
	}
}
