package main

import (
	"fmt"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/task"
	"github.com/hrdkgmz/cacheSync/taskHandle"
)

func startRefresh(pool *taskHandle.WorkPool)error{
	fmt.Println("开始分发数据库全量同步任务...")
	for k,_ := range global.GetHashInfos(){
		pool.Do(task.NewRefreshTask(k))
	}
	err:=pool.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库全量数据同步任务已全部完成")
	return nil
}