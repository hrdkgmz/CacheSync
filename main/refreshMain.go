package main

import (
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/task"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"sync"
)

func startRefresh() error {
	var wg sync.WaitGroup
	log.Info("开始分发数据库全量同步任务...")
	pool := taskHandle.GetInstance()
	wg.Add(len(global.GetHashInfos()))
	for k, _ := range global.GetHashInfos() {
		pool.Do(task.NewRefreshTask(k, &wg))
	}
	wg.Wait()
	log.Info("数据库全量数据同步任务已全部完成")
	return nil
}
