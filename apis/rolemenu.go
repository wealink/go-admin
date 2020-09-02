package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取菜单树
func GetTreeRoleMenus(c *gin.Context) {
	var menu models.Menu
	var rolemenu models.RoleMenu
	rolemenu.Roleid = com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	data["menus"] = menu.GetTreeMenus()
	data["checkedKeys"] = rolemenu.GetTreeRoleMenus()
	util.Response(c, e.Code_200, data, e.GetSuccess)
}

//获取菜单角色列表
func GetRoleMenus(c *gin.Context) {
	var rolemenu models.RoleMenu
	util.Response(c, e.Code_200, rolemenu.GetRoleMenus(), e.GetSuccess)
}
