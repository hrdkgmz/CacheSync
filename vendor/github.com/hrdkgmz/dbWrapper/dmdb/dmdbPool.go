package dmdb

import (
	"database/sql"
	"fmt"
	log "github.com/cihub/seelog"
	"sync"
)

var DbMan *DbPool
var once sync.Once

type DbPool struct {
	Db *SQLConnPool
}

func GetInstance() *DbPool {
	once.Do(func() {
		DbMan = NewDbPool()
		if err := DbMan.Open(); err != nil {
			log.Error("达梦数据库连接池初始化失败", err.Error())
			panic("达梦数据库连接失败，程序退出")
		}
		log.Info("达梦数据库连接成功！")
	})
	return DbMan
}

var _host string = "172.28.62.30:5236"
var _username string = "SYSDBA"
var _password string = "SYSDBA"
var _maxOpens int = 2000
var _maxIdels int = 1000

func SetMysqlParas(host, user, password string, maxOpenConns, maxIdleConns int) {
	_host = host
	_username = user
	_password = password
	_maxOpens = maxOpenConns
	_maxIdels = maxIdleConns

}

func NewDbPool() *DbPool {
	db := &DbPool{
		Db: initMySQLPool(_host, _username, _password, _maxOpens, _maxIdels),
	}

	return db
}

type SQLConnPool struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	SQLDB          *sql.DB
}

func initMySQLPool(host, user, password string, maxOpenConns, maxIdleConns int) *SQLConnPool {
	dataSourceName := fmt.Sprintf("dm://%s:%s@%s", user, password, host)
	//fmt.Println(dataSourceName)
	db := &SQLConnPool{
		DriverName:     "dm",
		DataSourceName: dataSourceName,
		MaxOpenConns:   maxOpenConns,
		MaxIdleConns:   maxIdleConns,
	}

	return db
}
