package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.OK
		token := c.Query("token")
		if token == "" {
			code = e.ErrParamCode
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrTimeoutCode
				default:
					code = e.ErrTokenCode
				}
			}
		}

		if code != e.OK {
			utils.Err(c, code)
			c.Abort()
			return
		}
		c.Next()
	}
}
