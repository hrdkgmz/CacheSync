package task

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/cache"
	"github.com/hrdkgmz/cacheSync/db"
	"github.com/hrdkgmz/cacheSync/global"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"strconv"
	"strings"
)

func NewDulpTask(tb string, key string) taskHandle.TaskHandler {
	return func() error {
		fmt.Println("执行数据库全量同步任务，表：" + tb + "， 缓存Key：" + key)
		mysql := db.GetInstance()
		sqlStr := "select * from " + tb
		list, err := mysql.Query(sqlStr)
		if err != nil {
			return err
		}
		err = cacheTable(tb, list, key)
		if err != nil {
			return err
		}
		return nil
	}
}

func cacheTable(tb string, list []map[string]interface{}, key string) error {
	ss := strings.Split(key, ";")
	for _, val := range list {
		for _, s := range ss {
			var key string
			if strings.Index(s, "&") >= 0 {
				multiKeys := strings.Split(s, "&")
				var realKeys string
				for i := 0; i < len(multiKeys); i++ {
					mk := multiKeys[i]
					if i == len(multiKeys)-1 {
						str, err := toString(val[mk])
						if err != nil {
							return err
						}
						realKeys += str
					} else {
						str, err := toString(val[mk])
						if err != nil {
							return err
						}
						realKeys += str + "&"
					}
				}
				key = tb + ":" + s + ":" + realKeys
			} else {
				kk := val[s]
				str, err := toString(kk)
				if err != nil {
					return err
				}
				key = tb + ":" + s + ":" + str
			}
			_, err := cache.GetInstance().SetHashMap(key, val)
			if err != nil {
				return err
			}
		}
		err := specialCase(tb, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func specialCase(tb string, val map[string]interface{}) error {
	switch tb {
	case "b_peer_info":
		peer, err := toString(val["peer_name"])
		if err != nil {
			return err
		}
		org, err := toString(val["org_name"])
		if err != nil {
			return err
		}
		orgPeerMap := global.GetOrgPeerMap()
		if orgPeerMap[org] == nil {
			orgPeerMap[org] = make([]string, 0)
		}
		orgPeerMap[org] = append(orgPeerMap[org], peer)
		return nil

	case "b_channel_info":
		cp, err := toString(val["chan_peer"])
		if err != nil {
			return err
		}
		peers := strings.Split(cp, ";")
		chann, err := toString(val["chan_name"])
		if err != nil {
			return err
		}
		chanPeerMap := global.GetChanPeerMap()
		chanPeerMap[chann] = make([]string, 0)
		for _, v := range peers {
			chanPeerMap[chann] = append(chanPeerMap[chann], v)
		}
		return nil
	case "b_orderer_info":
		o, err := toString(val["ord_name"])
		if err != nil {
			return err
		}
		global.AppendOrderers(o)
		return nil
	case "b_peer_cc":
		peer, err := toString(val["peer_name"])
		if err != nil {
			return err
		}
		cc, err := toString(val["cc_id"])
		if err != nil {
			return err
		}
		ccPeerMap := global.GeCCPeerMap()
		peerCCMap := global.GetPeerCCMap()
		if ccPeerMap[cc] == nil {
			ccPeerMap[cc] = make([]string, 0)
		}
		ccPeerMap[cc] = append(ccPeerMap[cc], peer)
		if peerCCMap[peer] == nil {
			peerCCMap[peer] = make([]string, 0)
		}
		peerCCMap[peer] = append(peerCCMap[peer], cc)

		for k, v := range ccPeerMap {
			ccPeerMap[k] = removeDuplicateElement(v)
		}
		for k, v := range peerCCMap {
			peerCCMap[k] = removeDuplicateElement(v)
		}
		return nil
	default:
		return nil
	}
}

func toString(v interface{}) (string, error) {
	switch v.(type) {

	case string:
		return v.(string), nil
	case int:
		return strconv.Itoa(v.(int)), nil
	case int64:
		return strconv.FormatInt(v.(int64), 10), nil
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64), nil
	default:
		log.Error("无法处理的数据类型")
		return "", nil
	}
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
