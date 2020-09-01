package util

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

// 响应数据处理数据处理
func Response(c *gin.Context, code int, data interface{}, msg string) {
	// 添加操作记录
	var optlog models.Optlog
	optlog.Username = GetUserName(c)
	optlog.Url = c.Request.URL.Path
	optlog.Method = c.Request.Method
	if code == e.Code_200 {
		optlog.Status = "0"
	} else {
		optlog.Status = "1"
	}
	optlog.Msg = msg
	optlog.Ipaddr = c.ClientIP()
	optlog.Optlocation = GetLocation(optlog.Ipaddr)
	optlog.Opttime = GetCurrentTime()
	optlog.AddOptLog()

	// 响应体
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	return
}
