package models

import "gin-example/global/orm"

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func (auth *Auth) Login() bool {
	var count int
	orm.Db.Table("go_user").Where("status = 0 and username = ? and password = ?", auth.Username, auth.Password).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
