package global

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"os"
)

type HashInfo struct {
	TbName string
	Keys   []string
	SetString bool
}

type SetInfo struct {
	TbName  string
	SetType SetType
	SetName []string
	Key     []string
	Member  []string
}

type SetType int

const (
	SetType_SingleKeySingleMember SetType = 1
	SetType_SingleKeyMultiMember  SetType = 2
	SetType_SingleMember          SetType = 3
	SetType_DoubleKeySingleMember SetType = 4
)

var (
	mapName   string = "syncMap"
	mapPath   string = "../config/"
	HashInfos map[string]*HashInfo
	SetInfos  map[string][]*SetInfo
)

func InitSyncMap() {
	v := viper.New()
	v.SetConfigName(mapName)
	v.AddConfigPath(mapPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		log.Error("缓存数据同步规则加载失败")
		os.Exit(1)
	}
	subV := v.Get("hash")
	for _, tb := range subV.([]interface{}) {
		if tbMap, ok := tb.(map[interface{}]interface{}); ok {
			tbName := tbMap["tbName"].(string)
			tbKeys := make([]string, 0)
			for _, v := range tbMap["keys"].([]interface{}) {
				tbKeys = append(tbKeys, v.(string))
			}
			isSetString := false
			if tbMap["setString"]!=nil{
				isSetString = tbMap["setString"].(bool)
			}
			tbInfo := HashInfo{tbName, tbKeys, isSetString}
			if HashInfos == nil {
				HashInfos = make(map[string]*HashInfo)
			}
			HashInfos[tbName] = &tbInfo
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
			sInfo := SetInfo{tbName, SetType(setType), setName, key, member}
			if SetInfos == nil {
				SetInfos = make(map[string][]*SetInfo)
			}
			if SetInfos[tbName] == nil {
				SetInfos[tbName] = make([]*SetInfo, 0)
			}
			SetInfos[tbName] = append(SetInfos[tbName], &sInfo)
		}
	}
	log.Info("缓存数据同步规则加载成功！")
}

func GetHashInfos() map[string]*HashInfo {
	return HashInfos
}

func GetSetInfos() map[string][]*SetInfo {
	return SetInfos
}
