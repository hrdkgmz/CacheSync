package global

type SetType int

const (
	SetType_SingleKeySingleMember SetType = 1
	SetType_SingleKeyMultiMember  SetType = 2
	SetType_SingleMember          SetType = 3
	SetType_DoubleKeySingleMember SetType = 4
)

type setInfo struct {
	tbName  string
	setType SetType
	setName []string
	key     []string
	member  []string
}

func newBlankSetInfo() *setInfo {
	return &setInfo{}
}

func newSetInfo(tbName string, setType SetType, setName []string, key []string, member []string) *setInfo {
	return &setInfo{tbName: tbName, setType: setType, setName: setName, key: key, member: member}
}

func (s *setInfo) TbName() string {
	return s.tbName
}

func (s *setInfo) SetTbName(tbName string) {
	s.tbName = tbName
}

func (s *setInfo) SetType() SetType {
	return s.setType
}

func (s *setInfo) SetSetType(setType SetType) {
	s.setType = setType
}

func (s *setInfo) SetName() []string {
	return s.setName
}

func (s *setInfo) SetSetName(setName []string) {
	s.setName = setName
}

func (s *setInfo) Key() []string {
	return s.key
}

func (s *setInfo) SetKey(key []string) {
	s.key = key
}

func (s *setInfo) Member() []string {
	return s.member
}

func (s *setInfo) SetMember(member []string) {
	s.member = member
}
