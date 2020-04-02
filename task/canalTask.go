package task

import (
	cProtocol "github.com/CanalClient/canal-go/protocol"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/util"
)

func NewDeleteTask(row *cProtocol.RowData, header *cProtocol.Header) func() error {
	return func() error {
		tb := header.GetTableName()
		tbInfo := global.GetSyncInfos()[tb]
		if tbInfo == nil {
			log.Info("未配置该数据表：" + tb + " 的缓存数据同步， binlog丢弃...")
			return nil
		}
		cols := row.GetBeforeColumns()
		colMap := make(map[string]interface{})
		for _, c := range cols {
			colMap[c.GetName()] = c.GetValue()
		}
		for _, key := range tbInfo.Keys() {
			rKey, err := util.BuildRedisKey(tb, key, colMap)
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().DelKey(rKey)
			if err != nil {
				return err
			}
			log.Info("已删除缓存数据，Key: "+rKey+" !")
		}
		if tbInfo.HasSpecial(){
			DeleteSpecial()
		}
		return nil
	}
}

func NewInsertTask(row *canal_protocol.RowData, header *canal_protocol.Header) func() error {
	return func() error {

	}
}

func NewUpdateTask(row *canal_protocol.RowData, header *canal_protocol.Header) func() error {
	return func() error {

	}
}
