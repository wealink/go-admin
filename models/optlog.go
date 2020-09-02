package models

import "gin-example/global/orm"

type Optlog struct {
	Base
	Username    string `json:"username"`    // 用户名
	Url         string `json:"url"`         // 请求url
	Method      string `json:"method"`      // 请求方法
	Status      string `json:"status"`      // 状态
	Msg         string `json:"msg"`         // 消息
	Ipaddr      string `json:"ipaddr"`      // IP地址
	Optlocation string `json:"optlocation"` // 归属地
	Opttime     int64  `json:"opttime"`     // 操作时间
}

func (Optlog) TableName() string {
	return "go_optlog"
}

//获取所有日志
func (log *Optlog) GetOptLogs(pageNum int, pageSize int) (logs []Optlog) {
	table := orm.Db.Table(log.TableName())
	if log.Username != "" {
		table = table.Where("username = ?", log.Username)
	}
	if log.Method != "" {
		table = table.Where("method = ?", log.Method)
	}
	if log.Status != "" {
		table = table.Where("status = ?", log.Status)
	}
	table.Order("Id DESC").Offset(pageNum).Limit(pageSize).Find(&logs)
	return
}

//获取所有日志总和
func (log *Optlog) GetOptLogsTotal() (count int) {
	table := orm.Db.Table(log.TableName())
	if log.Username != "" {
		table = table.Where("username = ?", log.Username)
	}
	if log.Method != "" {
		table = table.Where("method = ?", log.Method)
	}
	if log.Status != "" {
		table = table.Where("status = ?", log.Status)
	}
	table.Count(&count)
	return
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
