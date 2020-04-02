package util

import (
	log "github.com/cihub/seelog"
	"strconv"
	"strings"
)

func BuildRedisKey(tb string, key string, val map[string]interface{})(string,error){
	var rKey string
	if strings.Index(key, "&") >= 0 {
		multiKeys := strings.Split(key, "&")
		var realKeys string
		for i := 0; i < len(multiKeys); i++ {
			mk := multiKeys[i]
			if i == len(multiKeys)-1 {
				str, err := ToString(val[mk])
				if err != nil {
					return "", err
				}
				realKeys += str
			} else {
				str, err := ToString(val[mk])
				if err != nil {
					return "", err
				}
				realKeys += str + "&"
			}
		}
		rKey = tb + ":" + key + ":" + realKeys
	} else {
		kk := val[key]
		str, err := ToString(kk)
		if err != nil {
			return "",err
		}
		rKey = tb + ":" + key + ":" + str
	}
	return rKey,nil
}

func ToString(v interface{}) (string, error) {
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

func RemoveDuplicateElement(strs []string) []string {
	result := make([]string, 0, len(strs))
	temp := map[string]struct{}{}
	for _, item := range strs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
