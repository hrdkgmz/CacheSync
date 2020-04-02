package global

type syncInfo struct {
	tbName     string
	keys       []string
	hasSpecial bool
}

func (s *syncInfo) TbName() string {
	return s.tbName
}

func (s *syncInfo) SetTbName(tbName string) {
	s.tbName = tbName
}

func newBlankSyncInfo() *syncInfo {
	return &syncInfo{}
}

func newSyncInfo(tbName string, keys []string, hasSpecial bool) *syncInfo {
	return &syncInfo{tbName: tbName, keys: keys, hasSpecial: hasSpecial}
}

func (s *syncInfo) Keys() []string {
	return s.keys
}

func (s *syncInfo) SetKeys(keys []string) {
	s.keys = keys
}

func (s *syncInfo) HasSpecial() bool {
	return s.hasSpecial
}

func (s *syncInfo) SetHasSpecial(hasSpecial bool) {
	s.hasSpecial = hasSpecial
}
