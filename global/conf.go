package global

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"os"
)

type DbConf struct {
	Host string
	Database string
	Username string
	Password string
	Charset string
	MaxOpenConns int
	MaxIdleConns int
}

type CacheConf struct {
	Host string
	Password string
	Db int
	MaxOpenConns int
	MaxIdleConns int
}

type CanalConf struct {
	IP string
	Port int
	Username string
	Password string
	Destination string
	SoTimeOut int32
	IdleTimeOut int32
	Subscribe string
}

var (
	confName   string = "conf"
	confPath   string = "./config/"
	dbConfig *DbConf
	cacheConfig *CacheConf
	canalConfig *CanalConf
)

func InitConf() {
	v := viper.New()
	v.SetConfigName(confName)
	v.AddConfigPath(confPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		log.Error("配置文件加载失败")
		os.Exit(1)
	}
	dbConfig=new(DbConf)
	err= v.UnmarshalKey("mysql",dbConfig)
	if err!=nil{
		log.Error("数据库配置加载失败")
		os.Exit(1)
	}

	cacheConfig=new(CacheConf)
	err= v.UnmarshalKey("redis",cacheConfig)
	if err!=nil{
		log.Error("缓存配置加载失败")
		os.Exit(1)
	}

	canalConfig=new(CanalConf)
	err= v.UnmarshalKey("canal",canalConfig)
	if err!=nil{
		log.Error("canal配置加载失败")
		os.Exit(1)
	}
	log.Info("配置文件加载成功！")
}

func GetDbConf() *DbConf {
	return dbConfig
}

func GetCacheConf() *CacheConf{
	return cacheConfig
}

func GetCanalConf() *CanalConf{
	return canalConfig
}