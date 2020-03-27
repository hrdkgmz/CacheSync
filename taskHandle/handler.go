package taskHandle

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	once       sync.Once
	TaskPool   *WorkPool
	maxWorkers int           = 50
	timeOut    time.Duration = 30 * time.Second
)

func GetInstance() *WorkPool {
	once.Do(func() {
		TaskPool = NewPool(maxWorkers, timeOut)
		fmt.Println("任务处理线程池创建成功，最大goroutine数量：" + strconv.Itoa(maxWorkers) +
			", 任务超时时间：" + strconv.FormatFloat(timeOut.Seconds(), 'E', -1, 64) + "秒")
	})
	return TaskPool
}

func SetTaskPoolParam(max int, time time.Duration) {
	maxWorkers = max
	timeOut = time
}
