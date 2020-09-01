package models

import "gin-example/global/orm"

type Loginlog struct {
	Base
	Username		string `json:"username"`			// 用户名
	Status			string `json:"status"`				// 状态
	Ipaddr			string `json:"ipaddr"`              // IP地址
	Loginlocation   string `json:"loginlocation"`       // 归属地
	Browser 		string `json:"browser"`             // 浏览器
	Os 				string `json:"os"`                 	// 操作系统
	Logintime 		int64 `json:"logintime"`       		// 登录登出时间
	Msg 			string `json:"msg"`					// 响应消息
}


func (Loginlog) TableName() string {
	return "go_loginlog"
}

//获取所有日志
func (log *Loginlog) GetLoginLogs(pageNum int, pageSize int) (logs []Loginlog) {
	table := orm.Db.Table(log.TableName())
	if log.Ipaddr != "" {
		table = table.Where("ipaddr = ?", log.Ipaddr)
	}
	if log.Username != "" {
		table = table.Where("username = ?", log.Username)
	}
	if log.Status != "" {
		table = table.Where("status = ?", log.Status)
	}
	table.Order("Id").Offset(pageNum).Limit(pageSize).Find(&logs)
	return
}

//获取所有日志总和
func (log *Loginlog) GetLoginLogsTotal() (count int) {
	table := orm.Db.Table(log.TableName())
	if log.Ipaddr != "" {
		table = table.Where("ipaddr = ?", log.Ipaddr)
	}
	if log.Username != "" {
		table = table.Where("username = ?", log.Username)
	}
	if log.Status != "" {
		table = table.Where("status = ?", log.Status)
	}
    table.Count(&count)
	return
}
//添加登录登出记录
func (log *Loginlog) AddLoginLog() bool {
	orm.Db.Table(log.TableName()).Create(&Loginlog{
		Username: log.Username,
		Status: log.Status,
		Ipaddr: log.Ipaddr,
		Loginlocation: log.Loginlocation,
		Browser: log.Browser,
		Os: log.Os,
		Logintime: log.Logintime,
		Msg: log.Msg,
	})
	return true
}




