package global

import (
	"github.com/hrdkgmz/cacheSync/util"
	"strings"
)

var(
	orgPeerMap map[string][]string
	chanPeerMap map[string][]string
	orderers []string
	ccPeerMap map[string][]string
	peerCCMap map[string][]string
)

func GetOrgPeerMap() map[string][]string{
	if orgPeerMap==nil{
		orgPeerMap=make(map[string][]string)
	}
	return orgPeerMap
}

func GetChanPeerMap() map[string][]string{
	if chanPeerMap==nil{
		chanPeerMap=make(map[string][]string)
	}
	return chanPeerMap
}

func GetOrderers() []string{
	return orderers
}

func AppendOrderers(val string){
	if orderers==nil{
		orderers=make([]string,0)
	}
	orderers=append(orderers,val)
}

func GetPeerCCMap() map[string][]string{
	if peerCCMap==nil{
		peerCCMap=make(map[string][]string)
	}
	return peerCCMap
}

func GeCCPeerMap() map[string][]string{
	if ccPeerMap==nil{
		ccPeerMap=make(map[string][]string)
	}
	return ccPeerMap
}

func DulpSpecialCase(tb string, val map[string]interface{}) error {
	switch tb {
	case "b_peer_info":
		peer, err := util.ToString(val["peer_name"])
		if err != nil {
			return err
		}
		org, err := util.ToString(val["org_name"])
		if err != nil {
			return err
		}
		orgPeerMap := GetOrgPeerMap()
		if orgPeerMap[org] == nil {
			orgPeerMap[org] = make([]string, 0)
		}
		orgPeerMap[org] = append(orgPeerMap[org], peer)
		return nil

	case "b_channel_info":
		cp, err := util.ToString(val["chan_peer"])
		if err != nil {
			return err
		}
		peers := strings.Split(cp, ";")
		chann, err := util.ToString(val["chan_name"])
		if err != nil {
			return err
		}
		chanPeerMap := GetChanPeerMap()
		chanPeerMap[chann] = make([]string, 0)
		for _, v := range peers {
			chanPeerMap[chann] = append(chanPeerMap[chann], v)
		}
		return nil
	case "b_orderer_info":
		o, err := util.ToString(val["ord_name"])
		if err != nil {
			return err
		}
		AppendOrderers(o)
		return nil
	case "b_peer_cc":
		peer, err := util.ToString(val["peer_name"])
		if err != nil {
			return err
		}
		cc, err := util.ToString(val["cc_id"])
		if err != nil {
			return err
		}
		ccPeerMap := GeCCPeerMap()
		peerCCMap := GetPeerCCMap()
		if ccPeerMap[cc] == nil {
			ccPeerMap[cc] = make([]string, 0)
		}
		ccPeerMap[cc] = append(ccPeerMap[cc], peer)
		if peerCCMap[peer] == nil {
			peerCCMap[peer] = make([]string, 0)
		}
		peerCCMap[peer] = append(peerCCMap[peer], cc)

		for k, v := range ccPeerMap {
			ccPeerMap[k] = util.RemoveDuplicateElement(v)
		}
		for k, v := range peerCCMap {
			peerCCMap[k] = util.RemoveDuplicateElement(v)
		}
		return nil
	default:
		return nil
	}
}