package jwt

import (
	"gin-example/pkg/e"
	"gin-example/pkg/jwt"
	"github.com/gin-gonic/gin"

	//"net/http"
	jwt2 "github.com/dgrijalva/jwt-go"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			msg  string
			data interface{}
		)
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.Code_401
			msg = "token字段不存在，鉴权错误"
			c.JSON(code, gin.H{
				"captcha": code,
				"msg":     msg,
				"data":    data,
			})
			c.Abort()
			return
		} else {
			_, err := jwt.ParseToken(token)
			if err != nil {
				if ve, ok := err.(*jwt2.ValidationError); ok {
					if ve.Errors&(jwt2.ValidationErrorExpired|jwt2.ValidationErrorNotValidYet) != 0 {
						c.JSON(code, gin.H{
							"code": e.Code_50014,
							"msg":  e.Msg_50014,
							"data": data,
						})
						c.Abort()
						return
					}
				} else {
					code = e.Code_500
					msg = "token解析失败"
					c.JSON(code, gin.H{
						"code": code,
						"msg":  msg,
						"data": data,
					})
					c.Abort()
					return
				}
			}
		}
		claims, _ := jwt.ParseToken(token)
		c.Set("claims", claims)
		c.Next()
	}
}
