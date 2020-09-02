package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetOptLogs(c *gin.Context) {
	var log models.Optlog
	//log.Ipaddr = c.Query("ipaddr")
	//log.Username = c.Query("username")
	//log.Status = c.Query("status")
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = log.GetOptLogs(pageIndex, pageSize)
	data["total"] = log.GetOptLogsTotal()
	data["pageIndex"] = pageIndex
	data["pageSize"] = pageSize
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data": data,
	})
}
