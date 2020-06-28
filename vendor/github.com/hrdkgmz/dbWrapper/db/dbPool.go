package db

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
			log.Error("数据库连接池初始化失败", err.Error())
			panic("数据库连接失败，程序退出")
		}
		log.Info("数据库连接成功！")
	})
	return DbMan
}

var _host string = "192.168.110.164:3306"
var _database string = "bctest"
var _username string = "root"
var _password string = "Gmz_891019_ws"
var _charset string = "utf8"
var _maxOpens int = 2000
var _maxIdels int = 1000

func SetMysqlParas(host, database, user, password, charset string, maxOpenConns, maxIdleConns int) {
	_host = host
	_database = database
	_username = user
	_password = password
	_charset = charset
	_maxOpens = maxOpenConns
	_maxIdels = maxIdleConns

}

func NewDbPool() *DbPool {
	db := &DbPool{
		Db: initMySQLPool(_host, _database, _username, _password, _charset, _maxOpens, _maxIdels),
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

func initMySQLPool(host, database, user, password, charset string, maxOpenConns, maxIdleConns int) *SQLConnPool {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", user, password, host, database, charset)
	//fmt.Println(dataSourceName)
	db := &SQLConnPool{
		DriverName:     "mysql",
		DataSourceName: dataSourceName,
		MaxOpenConns:   maxOpenConns,
		MaxIdleConns:   maxIdleConns,
	}

	return db
}
