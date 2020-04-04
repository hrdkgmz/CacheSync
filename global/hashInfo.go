package global

type hashInfo struct {
	tbName     string
	keys       []string
}

func newBlankHashInfo() *hashInfo {
	return &hashInfo{}
}

func newHashInfo(tbName string, keys []string) *hashInfo {
	return &hashInfo{tbName: tbName, keys: keys}
}

func (s *hashInfo) TbName() string {
	return s.tbName
}

func (s *hashInfo) SetTbName(tbName string) {
	s.tbName = tbName
}

func (s *hashInfo) Keys() []string {
	return s.keys
}

func (s *hashInfo) SetKeys(keys []string) {
	s.keys = keys
}
