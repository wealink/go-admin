package permission

import (
	"gin-example/pkg/e"
	"gin-example/pkg/jwt"
	"gin-example/pkg/permissions"
	"github.com/gin-gonic/gin"
)

//权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("claims")
		v := data.(*jwt.Claims)
		//fmt.Println(v.Role_Name)
		casbin, err := permissions.Casbin()
		if err != nil {
			code := e.Code_500
			msg := e.Msg_500
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
			c.Abort()
			return
		} else {
			res, _ := casbin.Enforce(v.Rolename, c.Request.URL.Path, c.Request.Method)
			// fmt.Println(v.Rolename, c.Request.URL.Path, c.Request.Method, res)
			if res == false {
				code := e.Code_403
				msg := e.Msg_403
				c.JSON(code, gin.H{
					"code": code,
					"msg":  msg,
					"data": data,
				})
				c.Abort()
				return
			} else {
				c.Next()
			}
		}
	}
}
