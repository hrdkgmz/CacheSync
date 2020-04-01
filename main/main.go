package main

import (
	"fmt"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/db"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"time"
)

func main() {
	db.GetInstance()
	cache.GetInstance()

	taskHandle.SetTaskPoolParam(10, 30*time.Second)
	taskHandler := taskHandle.GetInstance()


	startDulp(taskHandler)

	for{
		time.Sleep(5 * time.Second)
		fmt.Printf("等待任务中...")
	}

	//go func() {
	//	for {
	//		time.Sleep(30 * time.Millisecond)
	//		taskHandler.Do(func() error {
	//			fmt.Println("任务执行成功")
	//			return nil
	//		})
	//	}
	//}()
	//var cnt int
	//for {
	//	time.Sleep(1 * time.Second)
	//	cnt++
	//	fmt.Printf("等待任务中，已等待%v秒\n", cnt)
	//
	//}
}
