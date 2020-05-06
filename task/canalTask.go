package task

import (
	cProtocol "github.com/CanalClient/canal-go/protocol"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/util"
	"strings"
)

func NewDeleteTask(row *cProtocol.RowData, header *cProtocol.Header) func() error {
	return func() error {
		tb := header.GetTableName()
		tbInfo := global.GetHashInfos()[tb]
		if tbInfo == nil {
			log.Info("未配置该数据表：" + tb + " 的缓存数据同步， binlog丢弃...")
			return nil
		}
		cols := row.GetBeforeColumns()
		colMap := make(map[string]interface{})
		for _, c := range cols {
			colMap[c.GetName()] = c.GetValue()
		}
		for _, key := range tbInfo.Keys {
			rKey, err := util.BuildRedisKey(tb, key, colMap)
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().DelKey(rKey)
			if err != nil {
				return err
			}
			log.Info("已删除缓存数据，Key: " + rKey)
		}
		if global.GetSetInfos()[tb] != nil {
			err := DeleteSetMember(tb, colMap)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func NewInsertTask(row *cProtocol.RowData, header *cProtocol.Header) func() error {
	return func() error {
		tb := header.GetTableName()
		tbInfo := global.GetHashInfos()[tb]
		if tbInfo == nil {
			log.Info("未配置该数据表：" + tb + " 的缓存数据同步， binlog丢弃...")
			return nil
		}
		cols := row.GetAfterColumns()
		colMap := make(map[string]interface{})
		for _, c := range cols {
			colMap[c.GetName()] = c.GetValue()
		}
		for _, key := range tbInfo.Keys {
			rKey, err := util.BuildRedisKey(tb, key, colMap)
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().SetHashMap(rKey, colMap)
			if err != nil {
				return err
			}
			log.Info("已新增缓存数据，Key: " + rKey)
		}
		if global.GetSetInfos()[tb] != nil {
			err := InsertSetMember(tb, colMap)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func NewUpdateTask(row *cProtocol.RowData, header *cProtocol.Header) func() error {
	return func() error {
		tb := header.GetTableName()
		tbInfo := global.GetHashInfos()[tb]
		if tbInfo == nil {
			log.Info("未配置该数据表：" + tb + " 的缓存数据同步， binlog丢弃...")
			return nil
		}

		beforeCols := row.GetBeforeColumns()
		bColMap := make(map[string]interface{})
		for _, c := range beforeCols {
			bColMap[c.GetName()] = c.GetValue()
		}
		afterCols := row.GetAfterColumns()
		aColMap := make(map[string]interface{})
		for _, c := range afterCols {
			aColMap[c.GetName()] = c.GetValue()
		}

		log.Info("开始更新缓存数据，删除旧key数据，新增新key数据")
		err := NewDeleteTask(row, header)()
		if err != nil {
			return err
		}
		err = NewInsertTask(row, header)()
		if err != nil {
			return err
		}

		if global.GetSetInfos()[tb] != nil {
			err := UpdateSetMember(tb, bColMap, aColMap)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func DeleteSetMember(tb string, val map[string]interface{}) error {
	setInfos := global.GetSetInfos()[tb]
	for _, setInfo := range setInfos {
		switch setInfo.SetType {
		case global.SetType_SingleKeySingleMember:
			mem, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			key, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().RemoveFrmSet(setInfo.SetName[0]+":"+key, mem)
			if err != nil {
				return err
			}

		case global.SetType_SingleKeyMultiMember:
			key, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().DelKey(setInfo.SetName[0] + ":" + key)
			if err != nil {
				return err
			}

		case global.SetType_SingleMember:
			mem, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().RemoveFrmSet(setInfo.SetName[0], mem)
			if err != nil {
				return err
			}

		case global.SetType_DoubleKeySingleMember:
			key1, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			mem1, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().RemoveFrmSet(setInfo.SetName[0]+":"+key1, mem1)
			if err != nil {
				return err
			}
			key2, err := util.ToString(val[setInfo.Key[1]])
			if err != nil {
				return err
			}
			mem2, err := util.ToString(val[setInfo.Member[1]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().RemoveFrmSet(setInfo.SetName[1]+":"+key2, mem2)
			if err != nil {
				return err
			}

		default:
			return nil
		}
	}
	return nil
}

func InsertSetMember(tb string, val map[string]interface{}) error {
	setInfos := global.GetSetInfos()[tb]
	for _, setInfo := range setInfos {
		switch setInfo.SetType {
		case global.SetType_SingleKeySingleMember:
			mem, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			key, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().AddToSet(setInfo.SetName[0]+":"+key, mem)
			if err != nil {
				return err
			}

		case global.SetType_SingleKeyMultiMember:
			mem, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			mems := strings.Split(mem, ";")
			key, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			for _, p := range mems {
				_, err = cache.GetInstance().AddToSet(setInfo.SetName[0]+":"+key, p)
				if err != nil {
					return err
				}
			}

		case global.SetType_SingleMember:
			mem, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().AddToSet(setInfo.SetName[0], mem)
			if err != nil {
				return err
			}

		case global.SetType_DoubleKeySingleMember:
			key1, err := util.ToString(val[setInfo.Key[0]])
			if err != nil {
				return err
			}
			mem1, err := util.ToString(val[setInfo.Member[0]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().AddToSet(setInfo.SetName[0]+":"+key1, mem1)
			if err != nil {
				return err
			}
			key2, err := util.ToString(val[setInfo.Key[1]])
			if err != nil {
				return err
			}
			mem2, err := util.ToString(val[setInfo.Member[1]])
			if err != nil {
				return err
			}
			_, err = cache.GetInstance().AddToSet(setInfo.SetName[1]+":"+key2, mem2)
			if err != nil {
				return err
			}

		default:
			return nil
		}
	}
	return nil
}

func UpdateSetMember(tb string, oldVal map[string]interface{}, newVal map[string]interface{}) error {
	err := DeleteSetMember(tb, oldVal)
	if err != nil {
		return err
	}
	err = InsertSetMember(tb, newVal)
	if err != nil {
		return err
	}
	return nil
}
