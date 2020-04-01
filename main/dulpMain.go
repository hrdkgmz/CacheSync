package main

import (
	"fmt"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/task"
	"github.com/hrdkgmz/cacheSync/taskHandle"
)

func startDulp(pool *taskHandle.WorkPool)error{
	keyMap := make(map[string]string)
	////////
	keyMap["b_org_info"]="org_name"
	keyMap["b_peer_info"]="peer_name"
	keyMap["b_user_info"] = "fab_user;sys_user"
	keyMap["b_orderer_info"]="ord_name"
	keyMap["b_channel_info"]="chan_name"
	keyMap["b_cc_info"]="cc_id"
	keyMap["b_peer_cc"]="peer_name&cc_id"
	keyMap["b_channel_info"]="chan_name"

	fmt.Println("开始分发数据库全量同步任务...")
	for k,v := range keyMap{
		pool.Do(task.NewDulpTask(k,v))
	}
	err:=pool.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库全量数据同步任务已全部完成")
	startHandleSepcialCase()
	return nil
}


func startHandleSepcialCase(){
	fmt.Println("开始数据库全量同步特殊处理任务...")

	for k, v := range global.GetOrgPeerMap() {
		cache.GetInstance().AddtoSet("org_peers:"+k, v...)
	}
	for k, v := range global.GetChanPeerMap() {
		cache.GetInstance().AddtoSet("chan_peers:"+k, v...)
	}
	cache.GetInstance().AddtoSet("orderers", global.GetOrderers()...)
	for k, v := range global.GetPeerCCMap() {
		cache.GetInstance().AddtoSet("peer_ccs:"+k, v...)
	}
	for k, v := range global.GeCCPeerMap() {
		cache.GetInstance().AddtoSet("cc_peers:"+k, v...)
	}
	fmt.Println("数据库全量同步特殊处理任务完成")
}