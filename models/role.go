package models

import (
	"gin-example/global/orm"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	Base
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type PostRole struct {
	Role
	Menuids []int `json:"menuids"`
}
func (Role) TableName() string {
	return "go_role"
}


//通过角色id查询角色是否存在
func (role *Role) ExistRoleByID() bool {
	orm.Db.Table(role.TableName()).Select("id").Where("id = ?", role.Id).First(&role)

	if role.Id > 0 {
		return true
	}

	return false
}

//获取所有角色
func (role *Role) GetRoles(pageNum int, pageSize int) (roles []Role) {
	table := orm.Db.Table(role.TableName())
	if role.Name != "" {
		table = table.Where("name = ?", role.Name)
	}
	if role.Status != "" {
		table = table.Where("status = ?", role.Status)
	}
	table.Offset(pageNum).Limit(pageSize).Find(&roles)
	return
}

//通过id获取角色信息
func (role *Role) GetRole() (roles Role) {
	table := orm.Db.Table(role.TableName())
	table.Where("id = ?", role.Id).Find(&roles)
	return
}

//获取所有角色总和
func (role *Role) GetRolesTotal() (count int) {
	table := orm.Db.Table(role.TableName())
	if role.Name != "" {
		table = table.Where("name = ?", role.Name)
	}
	if role.Status != "" {
		table = table.Where("status = ?", role.Status)
	}
	table.Count(&count)
	return
}

//添加角色
func (role *PostRole) AddRole() bool {
	// 判断角色是否冲突
	var count int
	var roleid []int
	orm.Db.Table(role.TableName()).Where("name = ?", role.Name).Count(&count)
	if count > 0 {
		return false
	}
	//先创建基本角色
	orm.Db.Table(role.TableName()).Create(&Role{
		Name:   role.Name,
		Status: role.Status,
	})
	orm.Db.Table(role.TableName()).Select("LAST_INSERT_ID() as id ").Where("name = ?", role.Name).Pluck("id", &roleid)
	//添加角色菜单关联数据
	if len(roleid) > 0 {
		var rolemenu RoleMenu
		for i := 0; i < len(role.Menuids); i++ {
			rolemenu.AddRoleMenu(roleid[0], role.Menuids[i])
		}
	}
	//添加菜单
	return true
}

//编辑角色
func (role *PostRole) EditRole() bool {
	orm.Db.Table(role.TableName()).Model(&Role{}).Where("id = ?", role.Id).Update(&role)
	var rolemenu RoleMenu
	rolemenu.DeleteRoleMenu(role.Id)
	for i := 0; i < len(role.Menuids); i++ {
		rolemenu.AddRoleMenu(role.Id, role.Menuids[i])
	}
	return true
}

//删除用户
func (role *Role) DeleteRole(id int) {
	orm.Db.Table(role.TableName()).Where("id = ?", id).Delete(Role{})
}

//创建时间和修改时间更新
func (role *Role) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (role *Role) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
