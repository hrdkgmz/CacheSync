package cache

import (
	log "github.com/cihub/seelog"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var once sync.Once
var RdsMan *RedisPool

type RedisPool struct {
	Pool *redis.Pool
}

func GetInstance() *RedisPool {
	once.Do(func() {
		RdsMan = NewInstance()
		if _, err := RdsMan.Do("PING"); err != nil {
			log.Error("Redis连接池初始化失败", err.Error())
			panic("Redis连接失败，程序退出")
		}
		log.Info("Redis连接成功！")
	})
	return RdsMan
}

var _host string = "192.168.110.164:6379"
var _pass string = ""
var _db int = 0
var _maxOpens int = 10032
var _maxIdels int = 300

func SetRedisParas(host, password string, database, maxOpenConns, maxIdleConns int) {
	_host = host
	_pass = password
	_db = database
	_maxOpens = maxOpenConns
	_maxIdels = maxIdleConns

}

func NewInstance() *RedisPool {
	red := &redis.Pool{
		MaxActive:   _maxOpens,
		MaxIdle:     _maxIdels,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", _host)
			if err != nil {
				return nil, err
			}
			if len(_pass) > 0 {
				if _, err := c.Do("AUTH", _pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("select", _db); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	redPool := &RedisPool{Pool: red}

	return redPool
}
