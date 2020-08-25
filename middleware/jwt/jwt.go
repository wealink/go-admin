package jwt

import (
	"gin-example/pkg/e"
	"gin-example/pkg/jwt"
	"github.com/gin-gonic/gin"
	//"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			msg  string
			data interface{}
		)
		//token := c.Query("token")
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
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = e.Code_500
				msg = "token解析失败"
				c.JSON(code, gin.H{
					"captcha": code,
					"msg":     msg,
					"data":    data,
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(code, gin.H{
					"captcha": e.Code_50014,
					"msg":     e.Code_50014,
					"data":    data,
				})
				c.Abort()
				return
			}
		}
		claims, _ := jwt.ParseToken(token)
		c.Set("claims", claims)
		c.Next()
	}
}
