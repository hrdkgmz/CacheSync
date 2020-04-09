package main

import (
	"bufio"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/db"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"os"
	"time"
)

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("../config/log-config/info.xml")
	if err != nil {
		fmt.Println("parse info.xml error:", err)
		return
	}
	log.ReplaceLogger(logger) //初始化日志实例

	global.InitConf()    //加载配置文件
	global.InitSyncMap() //加载同步规则映射表

	initParams() //配置参数加载

	db.GetInstance()    //连接数据库
	cache.GetInstance() //连接缓存

	//初始化任务队列与任务线程池
	taskHandle.GetInstance()

	time.Sleep(1 * time.Second)

	fmt.Println("运行模式选择：\n1.执行全量数据同步，并启动增量数据同步监听...\n2.仅启动增量数据同步监听...\n3.仅执行全量同步任务...\n请输入选项:")
	reader := bufio.NewReader(os.Stdin)

Read:
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "" {
			continue
		}
		switch command {
		case "1":
			startRefresh() //执行全量同步任务
			fallthrough
		case "2":
			if global.CanalConfig == nil {
				fmt.Println("canal配置异常，无法启动binlog监听...")
				time.Sleep(5 * time.Second)
				os.Exit(1)
			}
			go StartFetchEvent() //连接Canal服务，启动binlog监听
			break Read
		case "3":
			startRefresh() //执行全量同步任务
			time.Sleep(5 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("***无法识别的选项，请重新输入***")
			fmt.Println("运行模式选择：\n1.执行全量数据同步，并启动增量数据同步监听...\n2.仅启动增量数据同步监听...\n3.仅执行全量同步任务...\n请输入选项:")
			continue
		}
	}
	loop() //死循环，保持线程

}

func initParams() {
	dbConfig := global.DbConfig
	cacheConfig := global.CacheConfig
	canalConfig := global.CanalConfig
	poolConfig := global.TPoolConfig
	if dbConfig == nil {
		fmt.Println("数据库配置加载结果为空")
		os.Exit(1)
	}
	if cacheConfig == nil {
		fmt.Println("缓存配置加载结果为空")
		os.Exit(1)
	}

	db.SetMysqlParas(dbConfig.Host,
		dbConfig.Database,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Charset,
		dbConfig.MaxOpenConns,
		dbConfig.MaxIdleConns)

	cache.SetRedisParas(cacheConfig.Host,
		cacheConfig.Password,
		cacheConfig.Db,
		cacheConfig.MaxOpenConns,
		cacheConfig.MaxIdleConns)

	if poolConfig != nil {
		taskHandle.SetTaskPoolParam(poolConfig.MaxThreads,
			time.Duration(poolConfig.TimeOut)*time.Second)
	}

	if canalConfig == nil {
		fmt.Println("canal配置加载结果为空，无法启动监听binlog，但可执行全量同步任务")
		return
	}
	SetCanalParameters(canalConfig.IP,
		canalConfig.Port,
		canalConfig.Username,
		canalConfig.Password,
		canalConfig.Destination,
		canalConfig.SoTimeOut,
		canalConfig.IdleTimeOut,
		canalConfig.Schema)
}

func loop() {
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "" {
			continue
		}
		if command == "info" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/info.xml")
			if err != nil {
				fmt.Println("parse info.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动info日志模式")
		} else if command == "debug" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/debug.xml")
			if err != nil {
				fmt.Println("parse debug.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动debug日志模式")
		} else if command == "error" {
			logger, err := log.LoggerFromConfigAsFile("../config/log-config/error.xml")
			if err != nil {
				fmt.Println("parse error.xml error:", err)
				continue
			}
			log.ReplaceLogger(logger)
			fmt.Println("启动error日志模式")
		} else {
			fmt.Println("无法识别的命令")
			continue
		}

	}
}
