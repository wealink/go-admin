package permissions

import (
	"gin-example/global/orm"
	"gin-example/pkg/config"
	log "gin-example/pkg/logging"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func Casbin() (*casbin.Enforcer, error) {
	conn := orm.Conn
	Apter, err := gormadapter.NewAdapter(config.DatabaseConfig.Dbtype, conn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("config/rbac_model.conf", Apter)
	if err != nil {
		return nil, err
	}
	if err := e.LoadPolicy(); err == nil {
		return e, err
	} else {
		log.Fatal("casbin rbac_model or policy init error, message: %v \r\n", err.Error())
		return nil, err
	}
}
