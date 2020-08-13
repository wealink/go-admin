package apis

import (
	"gin-example/pkg/captcha"
	"gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		code := e.Code_500
		msg := "获取验证码失败"
		data := ""
		c.JSON(code, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
	} else {
		code := e.Code_200
		msg := "获取验证码成功"
		data := b64s
		id := id
		c.JSON(code, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
			"id":   id,
		})
	}
}
