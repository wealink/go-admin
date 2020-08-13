package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取所有菜单
func GetMenus(c *gin.Context) {
	menu := models.Menu{}
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = menu.GetMenus(pageIndex, pageSize)
	data["total"] = menu.GetMenuTotal()
	code := e.Code_200
	c.JSON(http.StatusOK, gin.H{
		"captcha": code,
		"msg":     e.GetSuccess,
		"data":    data,
	})

}

//新增菜单
func AddMenu(c *gin.Context) {
	name := c.Query("name")
	type1 := c.Query("type")
	path := c.Query("path")
	method := c.Query("method")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("菜单名不能为空")
	valid.Required(type1, "type").Message("菜单类型不能为空")
	valid.Required(path, "path").Message("路径不能为空")
	valid.Required(method, "method").Message("请求方法类型不能为空")

	data := make(map[string]interface{})
	if valid.HasErrors() == false {
		menu := models.Menu{}
		data["name"] = name
		data["type"] = type1
		data["path"] = path
		data["method"] = method
		menu.AddMenu(data)
		c.JSON(http.StatusOK, gin.H{
			"captcha": e.Code_200,
			"data":    make(map[string]interface{}),
			"msg":     e.CreatedSuccess,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"captcha": e.Code_400,
			"data":    valid.Errors,
			"msg":     e.CreatedFail,
		})
		return
	}
}

//编辑菜单
//编辑用户
func EditMenu(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	type1 := c.Query("type")
	path := c.Query("path")
	method := c.Query("method")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(name, "name").Message("菜单名不能为空")
	valid.Required(type1, "type").Message("菜单类型不能为空")
	valid.Required(path, "path").Message("路径不能为空")
	valid.Required(method, "method").Message("请求方法类型不能为空")

	data := make(map[string]interface{})
	if valid.HasErrors() == false {
		menu := models.Menu{}
		if menu.ExistMenuByID(id) {
			data["name"] = name
			menu.EditMenu(id, data)
			c.JSON(http.StatusOK, gin.H{
				"captcha": e.Code_200,
				"data":    make(map[string]interface{}),
				"msg":     e.UpdatedSuccess,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"captcha": e.Code_404,
				"data":    valid.Errors,
				"msg":     e.Msg_404,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"captcha": e.Code_400,
			"data":    valid.Errors,
			"msg":     e.UpdatedFail,
		})
		return
	}
}

//删除菜单
func DeleteMenu(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() == false {
		menu := models.Menu{}
		if menu.ExistMenuByID(id) {
			menu.DeleteMenu(id)
			c.JSON(http.StatusOK, gin.H{
				"captcha": e.Code_200,
				"data":    make(map[string]interface{}),
				"msg":     e.DeletedSuccess,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"captcha": e.Code_404,
				"data":    make(map[string]interface{}),
				"msg":     e.Msg_404,
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"captcha": e.Code_400,
			"data":    valid.Errors,
			"msg":     e.DeletedFail,
		})
		return
	}

}
