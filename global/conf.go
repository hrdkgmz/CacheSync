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
	Schema string
}

type TaskPoolConf struct {
	MaxThreads int
	TimeOut int
}

var (
	confName    string = "conf"
	confPath    string = "../config/"
	DbConfig    *DbConf
	CacheConfig *CacheConf
	CanalConfig *CanalConf
	TPoolConfig *TaskPoolConf
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
	DbConfig=new(DbConf)
	err= v.UnmarshalKey("mysql",DbConfig)
	if err!=nil{
		log.Error("数据库配置加载失败")
		os.Exit(1)
	}

	CacheConfig =new(CacheConf)
	err= v.UnmarshalKey("redis", CacheConfig)
	if err!=nil{
		log.Error("缓存配置加载失败")
		os.Exit(1)
	}

	CanalConfig =new(CanalConf)
	err= v.UnmarshalKey("canal", CanalConfig)
	if err!=nil{
		log.Error("canal配置加载失败")
		os.Exit(1)
	}

	TPoolConfig =new(TaskPoolConf)
	err= v.UnmarshalKey("taskpool", TPoolConfig)
	if err!=nil{
		log.Error("任务线程池配置加载失败")
		os.Exit(1)
	}
	log.Info("配置文件加载成功！")
}