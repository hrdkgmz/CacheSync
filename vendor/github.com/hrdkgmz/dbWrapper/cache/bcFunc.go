package cache

import (
	"errors"
	"strconv"
	"strings"
)

const (
	archiveDb int = 0
	taskDb    int = 1
)

func GetFabUserBySysUser(sysUser string) (string, error) {
	fabUser, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:sys_user:"+sysUser, "fab_user")
	return fabUser, err
}

func GetOrgByFabUser(fabUser string) (string, error) {
	org, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:fab_user:"+fabUser, "org_name")
	return org, err
}

func GetOrgBySysUser(sysUser string) (string, error) {
	org, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_user_info:hash:sys_user:"+sysUser, "org_name")
	return org, err
}

//组织下的节点
func GetPeerByOrg(org string) ([]string, error) {
	return getSetStringVals("org_peers:" + org)
}

//加入通道的节点 s_channel_peer
func GetPeerByChannel(channel string) ([]string, error) {
	return getSetStringVals("chan_peers:" + channel)
}

//一个组织下加入通道的节点
func GetCommonPeerByOrgChannel(org string, channel string) ([]string, error) {
	peer, err := GetInstance().GetSetInterWitchDb(archiveDb, "chan_peers:"+channel, "org_peers:"+org)
	strVals := make([]string, 0)
	for _, p := range peer {
		if str, ok := p.([]uint8); ok {
			strVals = append(strVals, string(str))
		}
	}
	if len(peer) != len(strVals) {
		return nil, errors.New("insufficient number of values")
	}
	return strVals, err
}

//查询链码安装到哪些peer上 s_peer_cc
func GetPeersByCCID(ccid string) ([]string, error) {
	return getSetStringVals("cc_peers:" + ccid)
}

//查询链码实例化在哪些通道上 s_channel_peer
func GetChansByCCID(ccid string) ([]string, error) {
	return getSetStringVals("cc_chans:" + ccid)
}

//通过ccid获取ccname和ccver
func GetCCInfoByCCID(ccid string) (string, string, error) {
	name, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccid, "cc_name")
	if err != nil {
		return name, "", err
	}
	ver, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccid, "cc_ver")
	return name, ver, nil
}
func GetOrderNames() ([]string, error) {
	return getSetStringVals("orderers")
}

func GetChannelTxByChannel(channel string) (string, error) {
	channelTxPath, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_channel_info:hash:chan_name:"+channel, "channeltx_path")
	return channelTxPath, err
}

func GetCCPathByCCID(ccID string) (string, error) {
	ccPath, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+ccID, "cc_path")
	return ccPath, err
}

func GetChanDetailByName(name string) (map[string]string, error) {
	channel, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_channel_info:hash:chan_name:"+name)
	return channel, err
}

//获取所有通道名
func GetAllChanName() ([]string, error) {
	return getSetStringVals("channels")
}

//获取所有通道信息
func GetAllChanDetail() ([]map[string]string, error) {
	chanNames, err := GetAllChanName()
	if err != nil {
		return nil, err
	}
	chanDetails := make([]map[string]string, 0)
	for _, v := range chanNames {
		chanDetail, err := GetChanDetailByName(v)
		if err != nil {
			return nil, err
		}
		chanDetails = append(chanDetails, chanDetail)
	}
	return chanDetails, nil
}

//通过ccid得到全量信息 b_cc_info
func GetCCDetailById(id string) (map[string]string, error) {
	cc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_cc_info:hash:cc_id:"+id)
	return cc, err
}

func GetCCDetailByCCNameAndCCVer(ccName string, ccVer string) (map[string]string, error) {
	cc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_cc_info:hash:cc_name&cc_ver:"+ccName+"&"+ccVer)
	return cc, err
}

func GetAllCCId() ([]string, error) {
	return getSetStringVals("ccs")
}

//所有链码信息 b_cc_info
func GetAllCCDetail() ([]map[string]string, error) {
	ccIds, err := GetAllCCId()
	if err != nil {
		return nil, err
	}
	ccDetails := make([]map[string]string, 0)
	for _, v := range ccIds {
		ccDetail, err := GetCCDetailById(v)
		if err != nil {
			return nil, err
		}
		ccDetails = append(ccDetails, ccDetail)
	}
	return ccDetails, nil
}

//通过peername获得peer全量信息 b_peer_info
func GetPeerDetailByName(name string) (map[string]string, error) {
	peer, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_peer_info:hash:peer_fullname:"+name)
	return peer, err
}

//加入通道的一个节点的一条记录 s_channel_peer
func GetChanPeerDetail(channel string, peer string) (map[string]string, error) {
	chanPeer, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_channel_peer:hash:chan_name&peer_fullname:"+channel+"&"+peer)
	return chanPeer, err
}

//获取peer加入通道的时间 s_channel_peer
func GetPeerJoinTime(channel string, peer string) (string, error) {
	time, err := GetInstance().GetHashStringWitchDb(archiveDb, "s_channel_peer:hash:chan_name&peer_fullname:"+channel+"&"+peer, "join_time")
	return time, err
}

//在通道中已实例化的一个链码的一条记录 s_channel_cc
func GetChanCCDetail(channel string, ccid string) (map[string]string, error) {
	chanCc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_channel_cc:hash:chan_name&cc_id:"+channel+"&"+ccid)
	return chanCc, err
}

//节点安装在链码链码上的一条记录 s_peer_cc
func GetPeerCCDetail(peer string, ccid string) (map[string]string, error) {
	chanCc, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "s_peer_cc:hash:peer_fullname&cc_id:"+peer+"&"+ccid)
	return chanCc, err
}

func GetUserByOrg(org string) ([]string, error) {
	return getSetStringVals("org_sysusers:" + org)
}

func GetOrgDetailByOrgName(orgName string) (map[string]string, error) {
	org, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_org_info:hash:org_name:"+orgName)
	return org, err
}

func GetAllOrgDetail() ([]map[string]string, error) {
	orgKeys, err := GetInstance().KeysWitchDb(archiveDb, "b_org_info:hash:org_name:*")
	if err != nil {
		return nil, err
	}
	orgDetails := make([]map[string]string, 0)
	for _, v := range orgKeys {
		orgdetail, err := GetInstance().GetHashMapStringWitchDb(archiveDb, v)
		if err != nil {
			return nil, err
		}
		orgDetails = append(orgDetails, orgdetail)
	}
	return orgDetails, nil
}

func GetAllPeerDetail() ([]map[string]string, error) {
	pKeys, err := GetInstance().KeysWitchDb(archiveDb, "b_peer_info:hash:peer_fullname:*")
	if err != nil {
		return nil, err
	}
	pDetails := make([]map[string]string, 0)
	for _, v := range pKeys {
		pdetail, err := GetInstance().GetHashMapStringWitchDb(archiveDb, v)
		if err != nil {
			return nil, err
		}
		pDetails = append(pDetails, pdetail)
	}
	return pDetails, nil
}

func GetOrderDetailByName(ordName string) (map[string]string, error) {
	ord, err := GetInstance().GetHashMapStringWitchDb(archiveDb, "b_orderer_info:hash:ord_name:"+ordName)
	return ord, err
}

func GetAllOrderDetail() ([]map[string]string, error) {
	oKeys, err := GetInstance().KeysWitchDb(archiveDb, "b_orderer_info:hash:ord_name:*")
	if err != nil {
		return nil, err
	}
	oDetails := make([]map[string]string, 0)
	for _, v := range oKeys {
		odetail, err := GetInstance().GetHashMapStringWitchDb(archiveDb, v)
		if err != nil {
			return nil, err
		}
		oDetails = append(oDetails, odetail)
	}
	return oDetails, nil
}

//所有ccname b_cc_info
func GetAllCCNames() ([]string, error) {
	ccKeys, err := GetInstance().KeysWitchDb(archiveDb, "b_cc_info:hash:cc_name&cc_ver:*")
	if err != nil {
		return nil, err
	}
	ccNames := make([]string, 0)
	for _, v := range ccKeys {
		cc := strings.TrimPrefix(v, "b_cc_info:hash:cc_name&cc_ver:")
		ccNames = append(ccNames, strings.Split(cc, "&")[0])
	}
	return RemoveDuplicateElement(ccNames), nil
}

//一个ccname的所有ccver
func GetAllCCVersByCCName(ccName string) ([]string, error) {
	return getSetStringVals("ccname_vers:" + ccName)
}

//通过链码名获取历史版本全部信息 b_cc_info
func GetCCDetailsByCCName(ccName string) ([]map[string]string, error) {
	ccVers, err := GetAllCCVersByCCName(ccName)
	if err != nil {
		return nil, err
	}
	ccDetails := make([]map[string]string, 0)
	for _, v := range ccVers {
		ccDetail, err := GetCCDetailByCCNameAndCCVer(ccName, v)
		if err != nil {
			return nil, err
		}
		ccDetails = append(ccDetails, ccDetail)
	}
	return ccDetails, nil
}

//一个peer加入的所有通道 s_channel_peer
func GetChannelByPeer(peerFullname string) ([]string, error) {
	return getSetStringVals("peer_chans:" + peerFullname)
}

func GetAllPeerCCDetail() ([]map[string]string, error) {
	pcKeys, err := GetInstance().KeysWitchDb(archiveDb, "s_peer_cc:hash:peer_fullname&cc_id:*")
	if err != nil {
		return nil, err
	}
	pcDetails := make([]map[string]string, 0)
	for _, v := range pcKeys {
		pcDetail, err := GetInstance().GetHashMapStringWitchDb(archiveDb, v)
		if err != nil {
			return nil, err
		}
		pcDetails = append(pcDetails, pcDetail)
	}
	return pcDetails, nil
}

func GetPeerCCDetailsByCCID(ccid string) ([]map[string]string, error) {
	pcKeys, err := GetPeersByCCID(ccid)
	if err != nil {
		return nil, err
	}
	pcDetails := make([]map[string]string, 0)
	for _, v := range pcKeys {
		pcDetail, err := GetPeerCCDetail(v, ccid)
		if err != nil {
			return nil, err
		}
		pcDetails = append(pcDetails, pcDetail)
	}
	return pcDetails, nil
}

func GetAllChannelCCDetail() ([]map[string]string, error) {
	chanCCKeys, err := GetInstance().KeysWitchDb(archiveDb, "s_channel_cc:hash:chan_name&cc_id:*")
	if err != nil {
		return nil, err
	}
	chanCCDetails := make([]map[string]string, 0)
	for _, v := range chanCCKeys {
		chanCCDetail, err := GetInstance().GetHashMapStringWitchDb(archiveDb, v)
		if err != nil {
			return nil, err
		}
		chanCCDetails = append(chanCCDetails, chanCCDetail)
	}
	return chanCCDetails, nil
}

func GetChanCCDetailsByCCID(ccid string) ([]map[string]string, error) {
	pcKeys, err := GetChansByCCID(ccid)
	if err != nil {
		return nil, err
	}
	pcDetails := make([]map[string]string, 0)
	for _, v := range pcKeys {
		pcDetail, err := GetChanCCDetail(v, ccid)
		if err != nil {
			return nil, err
		}
		pcDetails = append(pcDetails, pcDetail)
	}
	return pcDetails, nil
}

func GetCCIDByCCNameAndVer(ccName string, ccVer string) (string, error) {
	ccId, err := GetInstance().GetHashStringWitchDb(archiveDb, "b_cc_info:hash:cc_name&cc_ver:"+ccName+"&"+ccVer, "cc_id")
	return ccId, err
}

func GetCCIDsByCCName(ccName string) ([]string, error) {
	ccKeys, err := GetInstance().KeysWitchDb(archiveDb, "b_cc_info:hash:cc_name&cc_ver:"+ccName+"*")
	if err != nil {
		return nil, err
	}
	ccIds := make([]string, 0)
	for _, v := range ccKeys {
		ccId, err := GetInstance().GetHashStringWitchDb(archiveDb, v, "cc_id")
		if err != nil {
			return nil, err
		}
		ccIds = append(ccIds, ccId)
	}
	return ccIds, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SetTaskDataHash(taskID int64, val map[string]interface{}, timeoutSecond int) error {
	_, err := GetInstance().SetHashAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDataHashAll(taskID int64) (vals map[string]string, err error) {
	return GetInstance().GetHashMapStringWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func GetTaskDataHashSingle(taskID int64, field string) (val string, err error) {
	return GetInstance().GetHashStringWitchDb(taskDb, strconv.FormatInt(taskID, 10), field)
}

func SetTaskDataString(taskID int64, val interface{}, timeoutSecond int) error {
	if _, err := GetInstance().SetStringAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond); err != nil {
		return err
	}
	return nil
}

func GetTaskDataString(taskID int64) (val string, err error) {
	return GetInstance().GetStringWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func SetTaskDataList(taskID int64, val []string, timeoutSecond int) error {
	_, err := GetInstance().SetListToTailAndExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), val, timeoutSecond)
	if err != nil {
		return err
	}
	return nil
}

func GetTaskDataListByRange(taskID int64, start int, stop int) (val []string, err error) {
	return GetInstance().GetListByRangeWitchDb(taskDb, strconv.FormatInt(taskID, 10), start, stop)
}

func GetTaskDataListLen(taskID int64) (len int, err error) {
	return GetInstance().GetListLenWitchDb(taskDb, strconv.FormatInt(taskID, 10))
}

func IsTaskDataTimeout(taskID int64) (bool, error) {
	ttl, err := GetInstance().GetTTLWitchDb(taskDb, strconv.FormatInt(taskID, 10))
	if err != nil {
		return false, err
	}
	if ttl >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func RenewTaskDataTimout(taskID int64, timeoutSecond int) error {
	if _, err := GetInstance().SetExpireWitchDb(taskDb, strconv.FormatInt(taskID, 10), timeoutSecond); err != nil {
		return err
	}
	return nil
}

func getSetStringVals(key string) ([]string, error) {
	vals, err := GetInstance().GetSetMembersWitchDb(archiveDb, key)
	strVals := make([]string, 0)
	for _, o := range vals {
		if str, ok := o.([]uint8); ok {
			strVals = append(strVals, string(str))
		}
	}
	if len(vals) != len(strVals) {
		return nil, errors.New("insufficient number of values")
	}
	return strVals, err
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
