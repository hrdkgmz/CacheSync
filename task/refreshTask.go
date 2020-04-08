package task

import (
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/db"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"github.com/hrdkgmz/cacheSync/util"
	"strconv"
	"strings"
	"sync"
)

func NewRefreshTask(tb string, wg *sync.WaitGroup) taskHandle.TaskHandler {
	return func() error {
		log.Info("执行数据库全量同步任务，表：" + tb)
		info := global.GetHashInfos()[tb]
		var keyString strings.Builder
		for _, key := range info.Keys {
			keyString.WriteString(key + " ")
		}
		log.Debug(tb + ", 包含缓存key:" + keyString.String())
		mysql := db.GetInstance()
		sqlStr := "select * from " + tb
		list, err := mysql.Query(sqlStr)
		log.Debug(tb + ", 数据查询成功，记录条数为:：" + strconv.Itoa(len(list)))
		if err != nil {
			return err
		}
		err = cacheTable(tb, list, info.Keys)
		if err != nil {
			return err
		}
		if wg != nil {
			wg.Done()
		}
		return nil
	}
}

func cacheTable(tb string, list []map[string]interface{}, keys []string) error {
	log.Info(tb + ", 数据开始写入缓存...")
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
			log.Debug(rKey + ", 数据写入成功！")
		}
		if global.GetSetInfos()[tb] != nil {
			err := InsertSetMember(tb, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
