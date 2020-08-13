package util

import (
	"gin-example/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) (int, int) {
	pageIndex, _ := com.StrTo(c.Query("pageIndex")).Int()
	pageSize, _ := com.StrTo(c.Query("pageSize")).Int()
	if pageIndex > 0 {
		pageIndex = (pageIndex - 1) * config.ApplicationConfig.PageSize
	}

	return pageIndex, pageSize
}
