package global

import (
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"strings"
)

var (
	confName  string = "syncConf"
	confPath  string = "./config/"
	syncInfos map[string]*syncInfo
)

func InitSyncInfos() {
	v := viper.New()
	v.SetConfigName(confName)
	v.AddConfigPath(confPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		log.Error("syncConf配置文件加载失败")
	}
	subV := v.Get("table")
	for _, tb := range subV.([]interface{}) {
		if tbMap, ok := tb.(map[interface{}]interface{}); ok {
			tbName := tbMap["tbName"].(string)
			tbKeys := strings.Split(tbMap["keys"].(string), ";")
			tbSpecial := tbMap["hasSpecial"].(bool)
			tbInfo := newSyncInfo(tbName, tbKeys, tbSpecial)
			if syncInfos == nil {
				syncInfos = make(map[string]*syncInfo)
			}
			syncInfos[tbName] = tbInfo
		}
	}
}

func GetSyncInfos() map[string]*syncInfo {
	return syncInfos
}
