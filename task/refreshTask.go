package task

import (
	"fmt"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/db"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"github.com/hrdkgmz/cacheSync/util"
	"strings"
)

func NewRefreshTask(tb string) taskHandle.TaskHandler {
	return func() error {
		fmt.Println("执行数据库全量同步任务，表：" + tb)
		info := global.GetHashInfos()[tb]
		var keyString strings.Builder
		for _, key := range info.Keys() {
			keyString.WriteString(key + " ")
		}
		fmt.Println("表：" + tb + "， 包含缓存key：" + keyString.String())
		mysql := db.GetInstance()
		sqlStr := "select * from " + tb
		list, err := mysql.Query(sqlStr)
		if err != nil {
			return err
		}
		err = cacheTable(tb, list, info.Keys())
		if err != nil {
			return err
		}
		return nil
	}
}

func cacheTable(tb string, list []map[string]interface{}, keys []string) error {
	for _, val := range list {
		for _, key := range keys {
			rKey, err := util.BuildRedisKey(tb, key, val)
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().SetHashMap(rKey, val)
			if err != nil {
				return err
			}
		}
		if global.GetSetInfos()[tb]!=nil {
			err := InsertSetMember(tb, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
