package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取所有菜单
func GetMenus(c *gin.Context) {
	var menu models.Menu
	menu.Type = c.Query("type")
	util.Response(c, e.Code_200, menu.GetMenus(), e.GetSuccess)
}

//获取菜单树
func GetTreeMenus(c *gin.Context) {
	var menu models.Menu
	util.Response(c, e.Code_200, menu.GetTreeMenus(), e.GetSuccess)
}

//通过id获取菜单信息
func GetMenu(c *gin.Context) {
	var menu models.Menu
	menu.Id = com.StrTo(c.Param("id")).MustInt()
	util.Response(c, e.Code_200, menu.GetMenu(), e.GetSuccess)
}

//新增目录或菜单
func AddMenu(c *gin.Context) {
	var menu models.Menu
	err := c.BindJSON(&menu)
	if err == nil {
		rs := menu.AddMenu()
		if rs == true {
			util.Response(c, e.Code_200, "", e.CreatedSuccess)
		} else {
			util.Response(c, e.Code_500, "", "角色已经存在!")
		}
	} else {
		util.Response(c, e.Code_400, err, e.Msg_400)
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
			util.Response(c, e.Code_200, "", e.UpdatedSuccess)
		} else {
			util.Response(c, e.Code_404, "", e.Msg_404)
		}
	} else {
		util.Response(c, e.Code_400, "", e.Msg_400)
	}
}

//删除菜单
func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	menu.Id = com.StrTo(c.Param("id")).MustInt()
	if menu.ExistMenuByID() {
		menu.DeleteMenu()
		util.Response(c, e.Code_200, "", e.DeletedSuccess)
	} else {
		util.Response(c, e.Code_404, "", e.Msg_404)
	}
}
