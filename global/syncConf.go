package global

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
)

var (
	confName  string = "syncConf"
	confPath  string = "./config/"
	hashInfos map[string]*hashInfo
	setInfos  map[string]*setInfo
)

func InitSyncInfos() {
	v := viper.New()
	v.SetConfigName(confName)
	v.AddConfigPath(confPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		log.Error("缓存数据同步规则加载失败")
	}
	subV := v.Get("hash")
	for _, tb := range subV.([]interface{}) {
		if tbMap, ok := tb.(map[interface{}]interface{}); ok {
			tbName := tbMap["tbName"].(string)
			tbKeys := make([]string, 0)
			for _, v := range tbMap["keys"].([]interface{}) {
				tbKeys = append(tbKeys, v.(string))
			}
			tbInfo := newHashInfo(tbName, tbKeys)
			if hashInfos == nil {
				hashInfos = make(map[string]*hashInfo)
			}
			hashInfos[tbName] = tbInfo
		}
	}
	subVV := v.Get("set")
	for _, tb := range subVV.([]interface{}) {
		if tbMap, ok := tb.(map[interface{}]interface{}); ok {
			tbName := tbMap["tbName"].(string)
			setType := tbMap["setType"].(int)
			setName := make([]string, 0)
			for _, v := range tbMap["setName"].([]interface{}) {
				setName = append(setName, v.(string))
			}
			var key []string
			if tbMap["key"] != nil {
				key = make([]string, 0)
				for _, v := range tbMap["key"].([]interface{}) {
					key = append(key, v.(string))
				}
			}
			member := make([]string, 0)
			for _, v := range tbMap["member"].([]interface{}) {
				member = append(member, v.(string))
			}
			sInfo := newSetInfo(tbName, SetType(setType), setName, key, member)
			if setInfos == nil {
				setInfos = make(map[string]*setInfo)
			}
			setInfos[tbName] = sInfo
		}
	}
	fmt.Println("缓存数据同步规则加载成功！")
}

func GetHashInfos() map[string]*hashInfo {
	return hashInfos
}

func GetSetInfos() map[string]*setInfo {
	return setInfos
}
