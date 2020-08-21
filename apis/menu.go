package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取所有菜单
func GetMenus(c *gin.Context) {
	var menu models.Menu
	menu.Type = c.Query("type")
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":     e.GetSuccess,
		"data":    menu.GetMenus(),
	})
}

//获取菜单树
func GetTreeMenus(c *gin.Context) {
	var menu models.Menu
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":     e.GetSuccess,
		"data":    menu.GetTreeMenus(),
	})
}

//通过id获取菜单信息
func GetMenu(c *gin.Context) {
	var menu models.Menu
	menu.Id = com.StrTo(c.Param("id")).MustInt()
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data": menu.GetMenu(),
	})
}

//新增目录或菜单
func AddMenu(c *gin.Context) {
	var menu models.Menu
	err := c.BindJSON(&menu)
	if err == nil {
		rs := menu.AddMenu()
		if rs == true {
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data": make(map[string]interface{}),
				"msg":  e.CreatedSuccess,
			})
			return
		} else {
			c.JSON(e.Code_500, gin.H{
				"code": e.Code_500,
				"data": nil,
				"msg":  "角色已经存在！！！",
			})
		}
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data": err,
			"msg":  e.Msg_400,
		})
		return
	}
}

//编辑角色
func EditMenu(c *gin.Context) {
	var menu models.Menu
	menu.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&menu)
	if err == nil {
		if menu.ExistMenuByID() {
			menu.EditMenu()
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data": make(map[string]interface{}),
				"msg":  e.UpdatedSuccess,
			})
			return
		} else {
			c.JSON(e.Code_404, gin.H{
				"code": e.Code_404,
				"data": make(map[string]interface{}),
				"msg":  e.Msg_404,
			})
			return
		}
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data": "",
			"msg":  e.Msg_400,
		})
		return
	}
}

//删除用户
func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	menu.Id = com.StrTo(c.Param("id")).MustInt()
	if menu.ExistMenuByID() {
		menu.DeleteMenu()
		c.JSON(e.Code_200, gin.H{
			"code": e.Code_200,
			"data": make(map[string]interface{}),
			"msg":  e.DeletedSuccess,
		})
		return
	} else {
		c.JSON(e.Code_404, gin.H{
			"code": e.Code_404,
			"data": make(map[string]interface{}),
			"msg":  e.Msg_404,
		})
		return
	}

}