package global

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"os"
)

var (
	confName   string = "conf"
	confPath   string = "./config/"
	dbConfig *dbConf
	cacheConfig *cacheConf
	canalConfig *canalConf
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
	dbConfig=new(dbConf)
	err= v.UnmarshalKey("mysql",dbConfig)
	if err!=nil{
		log.Error("数据库配置加载失败")
		os.Exit(1)
	}

	cacheConfig=new(cacheConf)
	err= v.UnmarshalKey("redis",cacheConfig)
	if err!=nil{
		log.Error("缓存配置加载失败")
		os.Exit(1)
	}

	canalConfig=new(canalConf)
	err= v.UnmarshalKey("redis",canalConfig)
	if err!=nil{
		log.Error("canal配置加载失败")
		os.Exit(1)
	}
	log.Info("配置文件加载成功！")
}

func GetDbConf() *dbConf{
	return dbConfig
}

func GetCacheConf() *cacheConf{
	return cacheConfig
}

func GetCanalConf() *canalConf{
	return canalConfig
}