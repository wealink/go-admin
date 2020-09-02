package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetLoginLogs(c *gin.Context) {
	var log models.Loginlog
	log.Ipaddr = c.Query("ipaddr")
	log.Username = c.Query("username")
	log.Status = c.Query("status")
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = log.GetLoginLogs(pageIndex, pageSize)
	data["total"] = log.GetLoginLogsTotal()
	data["pageIndex"] = pageIndex
	data["pageSize"] = pageSize
	util.Response(c, e.Code_200, data, e.GetSuccess)
}
