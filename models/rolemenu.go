package models

import (
	"fmt"
	"gin-example/global/orm"
	"github.com/chenhg5/collection"
)

type RoleMenu struct {
	Base
	Roleid int `json:"roleid"`
	Menuid int `json:"menuid"`
}

type MenuAddRoles struct {
	Menu
	Roles []string `json:"roles"`
}
type MenuViewAddRoles struct {
	Menu
	Children []MenuAddRoles `json:"children"`
	Roles []string `json:"roles"`
}

type CasbinRule struct {
	Ptype string `gorm:"column:p_type"`
	Rolename string `gorm:"column:v0"`
	Path string `gorm:"column:v1"`
	Method string `gorm:"column:v2"`
}
func (RoleMenu) TableName() string {
	return "go_role_menu"
}

//添加角色菜单关联
func (rolemenu *RoleMenu) AddRoleMenu(roleid,menuid int) bool {
	var menu Menu
	var role Role
	orm.Db.Table("go_menu").Where("id = ?",menuid).Find(&menu)
	orm.Db.Table("go_role").Where("id = ?",roleid).Find(&role)
	if menu.Type == "J" {
		orm.Db.Table("casbin_rule").Create(&CasbinRule{
			Ptype: "p",
			Rolename: role.Name,
			Path: menu.Path,
			Method: menu.Method,
		})
	}
	orm.Db.Table(rolemenu.TableName()).Create(&RoleMenu{
		Roleid: roleid,
		Menuid: menuid,
	})
	return true
}

//通过角色id查询菜单列表
func (rolemenu *RoleMenu) GetRoleMenu()(menuids []int) {
	var menu MenuView
	var tmp []int
	//获取目录ID列表
	menus :=menu.GetMenus()
	for i:=0;i<len(menus);i++{
		tmp=append(tmp, menus[i].Id)
		for k:=0;k<len(menus[i].Children);k++{
			tmp=append(tmp, menus[i].Children[k].Id)
		}
	}
	fmt.Println(tmp)
	//获取关联列表
	orm.Db.Table(rolemenu.TableName()).Select("menuid").Where("roleid = ?", rolemenu.Roleid).Pluck("menuid",&menuids)
	//剔除关联列表中半选中的目录ID
	if len(menuids)-len(tmp) >=0 {
		for j := len(menuids) - len(menus); j < len(menuids); j++ {
			if collection.Collect(tmp).Contains(menuids[j]) == true {
				menuids = append(menuids[:j], menuids[j+1:]...)
			}
		}
	}
	return
}

//获取菜单树并添加上roles字段，用于前端路由拼接
func (rolemenu *RoleMenu) GetRoleMenus()(menusparent []MenuViewAddRoles) {
	var menuschildren []MenuAddRoles
	orm.Db.Table("go_menu").Where("componenturl = 'Layout'").Find(&menusparent)
	for index := 0; index < len(menusparent); index++ {
		orm.Db.Table("go_role_menu").Select([]string{"go_role.name"}).Joins("left join go_role on go_role_menu.roleid=go_role.id").Where("go_role_menu.menuid = ? ", menusparent[index].Id).Pluck("go_role.name",&menusparent[index].Roles)
		fmt.Println(menusparent[index].Roles)
		orm.Db.Table("go_menu").Where("parentid = ? ", menusparent[index].Id).Find(&menuschildren)
		for index := 0; index < len(menuschildren); index++ {
			orm.Db.Table("go_role_menu").Select([]string{"go_role.name as roles"}).Joins("left join go_role on go_role_menu.roleid=go_role.id").Where("go_role_menu.menuid = ? ", menuschildren[index].Id).Pluck("go_role.name",&menuschildren[index].Roles)
		}
		menusparent[index].Children = menuschildren
	}
	return
}

//删除角色菜单关联
func (rolemenu *RoleMenu) DeleteRoleMenu(roleid int) {
	var role Role
	orm.Db.Table("go_role").Where("id = ?",roleid).Find(&role)
	orm.Db.Table("casbin_rule").Where("v0 = ?", role.Name).Delete(CasbinRule{})
	orm.Db.Table(rolemenu.TableName()).Where("roleid = ?", roleid).Delete(RoleMenu{})
}


