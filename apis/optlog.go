package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetOptLogs(c *gin.Context) {
	var log models.Optlog
	log.Username = c.Query("username")
	log.Method = c.Query("method")
	log.Status = c.Query("status")
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = log.GetOptLogs(pageIndex, pageSize)
	data["total"] = log.GetOptLogsTotal()
	data["pageIndex"] = pageIndex
	data["pageSize"] = pageSize
	util.Response(c, e.Code_200, data, e.GetSuccess)
}
