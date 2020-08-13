package database

import (
	"bytes"
	"gin-example/global/orm"
	"gin-example/pkg/config"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func InitDb() {
	var err error
	orm.Conn = GetConnect()
	orm.Db, err = Open(DbType, orm.Conn)
	if err != nil {
		log.Fatalf("%s connect error %v", DbType, err)
	} else {
		log.Printf("%s connect success!", DbType)
	}
	orm.Db.SingularTable(true)
	orm.Db.LogMode(true)
	orm.Db.DB().SetMaxIdleConns(10)
	orm.Db.DB().SetMaxOpenConns(30)
}

type Mysql struct{}

func Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

func GetConnect() string {

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}

func CloseDB() {
	defer orm.Db.Close()
}
