package models

import (
	"fmt"
	"gin-example/global/orm"
	"github.com/jinzhu/gorm"
	"time"
)

//type User struct {
//	Base
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Role_Id  int    `json:"role_id"`
//}

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	Roleid   int    `json:"roleid"`
	// Rolename string `json:"rolename" gorm:"column:name"`
}

type UserView struct {
	User
	Rolename string `json:"rolename" gorm:"column:name"`
}

func (User) TableName() string {
	return "go_user"
}

func (user *User) Login(username, password string) bool {
	orm.Db.Select("id").Where("username = ? and password = ?", username, password).First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

//通过用户查询角色名
func (user *User) GetRoleNameByUserName(username string) string {
	var userview UserView
	orm.Db.Table(user.TableName()).Select([]string{"go_role.name"}).Joins("left join go_role on go_user.roleid=go_role.id").Where("go_user.username = ?", username).Find(&userview)
	return userview.Rolename
}

//通过用户id查询用户是否存在
func (user *User) ExistUserByID() bool {
	orm.Db.Table(user.TableName()).Select("id").Where("id = ?", user.Id).First(&user)
	if user.Id > 0 {
		return true
	}

	return false
}

//获取所有用户
func (user *User) GetUsers(pageNum int, pageSize int) (users []UserView) {
	table := orm.Db.Table(user.TableName()).Select([]string{"go_user.*", "go_role.name"}).Joins("left join go_role on go_user.roleid=go_role.id")
	if user.Username != "" {
		table = table.Where("username = ?", user.Username)
	}
	if user.Phone != "" {
		table = table.Where("phone = ?", user.Phone)
	}
	if user.Status != "" {
		table = table.Where("status = ?", user.Status)
	}
	//table.Offset(pageNum).Limit(pageSize).Find(&UserView{})
	table.Order("Id").Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

//通过id获取角色信息
func (user *User) GetUser() (users User) {
	table := orm.Db.Table(user.TableName())
	if user.Username == "" {
		table.Where("id = ?", user.Id).Find(&users)
	} else {
		table.Where("username = ?", user.Username).Find(&users)
	}
	return
}

//获取所有用户总和
func (user *User) GetUserTotal() (count int) {
	table := orm.Db.Table(user.TableName())
	if user.Username != "" {
		table = table.Where("username = ?", user.Username)
	}
	if user.Phone != "" {
		table = table.Where("phone = ?", user.Phone)
	}
	if user.Status != "" {
		table = table.Where("status = ?", user.Status)
	}
	table.Count(&count)
	return
}

//添加用户
func (user *User) AddUser() bool {
	// 判断角色是否冲突
	var count int
	orm.Db.Table(user.TableName()).Where("name = ?", user.Username).Count(&count)
	if count > 0 {
		return false
	}
	orm.Db.Table(user.TableName()).Create(&User{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Status:   user.Status,
		Roleid:   user.Roleid,
	})
	return true
}

//编辑用户
func (user *User) EditUser() bool {
	orm.Db.Table(user.TableName()).Model(&User{}).Where("id = ?", user.Id).Update(&user)
	return true
}

//个人中心 上传用户头像
func (user *User) UploadAvatar() bool {
	orm.Db.Table(user.TableName()).Where("username = ?", user.Username).Update("avatar", user.Avatar)
	return true
}

//个人中心 更新用户密码
func (user *User) Pwd(oldpassword, newpassword string) bool {
	fmt.Println(oldpassword, newpassword)
	orm.Db.Table(user.TableName()).Select("id").Where("username = ? and password = ?", user.Username, oldpassword).First(&user)
	if user.Id > 0 {
		orm.Db.Table(user.TableName()).Model(&User{}).Where("id = ?", user.Id).Update("password", newpassword)
		return true
	} else {
		return false
	}
}

//个人中心 修改用户头像

//删除用户
func (user *User) DeleteUser() {
	orm.Db.Table(user.TableName()).Where("id = ?", user.Id).Delete(User{})
}

//重置用户密码
func (user *User) ResetUserPwd() {
	orm.Db.Table(user.TableName()).Where("id = ?", user.Id).Update("password", &user.Password)
}

//创建时间和修改时间更新
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
