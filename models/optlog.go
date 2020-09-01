package models

import "gin-example/global/orm"

type Optlog struct {
	Base
	Username    string `json:"username"`    // 用户名
	Url         string `json:"url"`         // 请求url
	Method      string `json:"method"`      //请求方法
	Status      string `json:"status"`      // 状态
	Msg         string `json:"msg"`         // 消息
	Ipaddr      string `json:"ipaddr"`      // IP地址
	Optlocation string `json:"optlocation"` // 归属地
	Opttime     int64  `json:"opttime"`     // 操作时间
}

func (Optlog) TableName() string {
	return "go_optlog"
}

//添加登录登出记录
func (log *Optlog) AddOptLog() bool {
	orm.Db.Table(log.TableName()).Create(&Optlog{
		Username:    log.Username,
		Url:         log.Url,
		Method:      log.Method,
		Status:      log.Status,
		Msg:         log.Msg,
		Ipaddr:      log.Ipaddr,
		Optlocation: log.Optlocation,
		Opttime:     log.Opttime,
	})
	return true
}
