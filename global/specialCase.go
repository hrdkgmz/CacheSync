package global


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