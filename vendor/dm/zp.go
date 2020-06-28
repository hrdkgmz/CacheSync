/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_662 = "7.6.0.0"

	Dm_build_663 = "7.0.0.9"

	Dm_build_664 = "8.0.0.73"

	Dm_build_665 = "7.1.2.128"

	Dm_build_666 = "7.1.5.144"

	Dm_build_667 = "7.1.6.123"

	Dm_build_668 = 32768 - 128

	Dm_build_669 int16 = 1

	Dm_build_670 int16 = 2

	Dm_build_671 int16 = 3

	Dm_build_672 int16 = 4

	Dm_build_673 int16 = 5

	Dm_build_674 int16 = 6

	Dm_build_675 int16 = 7

	Dm_build_676 int16 = 8

	Dm_build_677 int16 = 9

	Dm_build_678 int16 = 13

	Dm_build_679 int16 = 14

	Dm_build_680 int16 = 15

	Dm_build_681 int16 = 17

	Dm_build_682 int16 = 21

	Dm_build_683 int16 = 24

	Dm_build_684 int16 = 27

	Dm_build_685 int16 = 29

	Dm_build_686 int16 = 30

	Dm_build_687 int16 = 31

	Dm_build_688 int16 = 32

	Dm_build_689 int16 = 44

	Dm_build_690 int16 = 52

	Dm_build_691 int16 = 60

	Dm_build_692 int16 = 71

	Dm_build_693 int16 = 90

	Dm_build_694 int16 = 91

	Dm_build_695 int16 = 200

	Dm_build_696 = 64

	Dm_build_697 = 20

	Dm_build_698 = 0

	Dm_build_699 = 4

	Dm_build_700 = 6

	Dm_build_701 = 10

	Dm_build_702 = 14

	Dm_build_703 = 18

	Dm_build_704 = 19

	Dm_build_705 = 128

	Dm_build_706 = 256

	Dm_build_707 = 0xffff

	Dm_build_708 int32 = 2

	Dm_build_709 = -1

	Dm_build_710 uint16 = 0xFFFE

	Dm_build_711 uint16 = uint16(Dm_build_710 - 3)

	Dm_build_712 uint16 = Dm_build_710

	Dm_build_713 int32 = 0xFFFF

	Dm_build_714 int32 = 0x80

	Dm_build_715 byte = 0x60

	Dm_build_716 uint16 = uint16(Dm_build_712)

	Dm_build_717 uint16 = uint16(Dm_build_713)

	Dm_build_718 int16 = 0x00

	Dm_build_719 int16 = 0x03

	Dm_build_720 int32 = 0x80

	Dm_build_721 byte = 0

	Dm_build_722 byte = 1

	Dm_build_723 byte = 2

	Dm_build_724 byte = 3

	Dm_build_725 byte = 4

	Dm_build_726 byte = Dm_build_721

	Dm_build_727 int = 10

	Dm_build_728 int32 = 32

	Dm_build_729 int32 = 65536

	Dm_build_730 byte = 0

	Dm_build_731 byte = 1

	Dm_build_732 int32 = 0x00000000

	Dm_build_733 int32 = 0x00000020

	Dm_build_734 int32 = 0x00000040

	Dm_build_735 int32 = 0x00000FFF

	Dm_build_736 int32 = 0

	Dm_build_737 int32 = 1

	Dm_build_738 int32 = 2

	Dm_build_739 int32 = 3

	Dm_build_740 = 8192

	Dm_build_741 = 1

	Dm_build_742 = 2

	Dm_build_743 = 0

	Dm_build_744 = 0

	Dm_build_745 = 1

	Dm_build_746 = -1

	Dm_build_747 int16 = 0

	Dm_build_748 int16 = 1

	Dm_build_749 int16 = 2

	Dm_build_750 int16 = 3

	Dm_build_751 int16 = 4

	Dm_build_752 int16 = 127

	Dm_build_753 int16 = Dm_build_752 + 20

	Dm_build_754 int16 = Dm_build_752 + 21

	Dm_build_755 int16 = Dm_build_752 + 22

	Dm_build_756 int16 = Dm_build_752 + 24

	Dm_build_757 int16 = Dm_build_752 + 25

	Dm_build_758 int16 = Dm_build_752 + 26

	Dm_build_759 int16 = Dm_build_752 + 30

	Dm_build_760 int16 = Dm_build_752 + 31

	Dm_build_761 int16 = Dm_build_752 + 32

	Dm_build_762 int16 = Dm_build_752 + 33

	Dm_build_763 int16 = Dm_build_752 + 35

	Dm_build_764 int16 = Dm_build_752 + 38

	Dm_build_765 int16 = Dm_build_752 + 39

	Dm_build_766 int16 = Dm_build_752 + 51

	Dm_build_767 int16 = Dm_build_752 + 71

	Dm_build_768 int16 = Dm_build_752 + 124

	Dm_build_769 int16 = Dm_build_752 + 125

	Dm_build_770 int16 = Dm_build_752 + 126

	Dm_build_771 int16 = Dm_build_752 + 127

	Dm_build_772 int16 = Dm_build_752 + 128

	Dm_build_773 int16 = Dm_build_752 + 129

	Dm_build_774 byte = 0

	Dm_build_775 byte = 2

	Dm_build_776 = 2048

	Dm_build_777 = -1

	Dm_build_778 = 0

	Dm_build_779 = 16000

	Dm_build_780 = 32000

	Dm_build_781 = 0x00000000

	Dm_build_782 = 0x00000020

	Dm_build_783 = 0x00000040

	Dm_build_784 = 0x00000FFF
)

type dm_build_785 interface {
	dm_build_786()
	dm_build_787() error
	dm_build_788()
	dm_build_789(imsg dm_build_785) error
	dm_build_790() error
	dm_build_791() (interface{}, error)
	dm_build_792()
	dm_build_793(imsg dm_build_785) (interface{}, error)
	dm_build_794()
	dm_build_795() error
	dm_build_796() byte
	dm_build_797() int32
	dm_build_798(length int32)
	dm_build_799() int16
}

type dm_build_800 struct {
	dm_build_801 *dm_build_410

	dm_build_802 int16

	dm_build_803 int32

	dm_build_804 *DmStatement
}

func (dm_build_806 *dm_build_800) dm_build_805(dm_build_807 *dm_build_410, dm_build_808 int16) *dm_build_800 {
	dm_build_806.dm_build_801 = dm_build_807
	dm_build_806.dm_build_802 = dm_build_808
	return dm_build_806
}

func (dm_build_810 *dm_build_800) dm_build_809(dm_build_811 *dm_build_410, dm_build_812 int16, dm_build_813 *DmStatement) *dm_build_800 {
	dm_build_810.dm_build_805(dm_build_811, dm_build_812).dm_build_804 = dm_build_813
	return dm_build_810
}

func dm_build_814(dm_build_815 *dm_build_410, dm_build_816 int16) *dm_build_800 {
	return new(dm_build_800).dm_build_805(dm_build_815, dm_build_816)
}

func dm_build_817(dm_build_818 *dm_build_410, dm_build_819 int16, dm_build_820 *DmStatement) *dm_build_800 {
	return new(dm_build_800).dm_build_809(dm_build_818, dm_build_819, dm_build_820)
}

func (dm_build_822 *dm_build_800) dm_build_786() {
	dm_build_822.dm_build_801.dm_build_413.Dm_build_92(0)
	dm_build_822.dm_build_801.dm_build_413.Dm_build_103(Dm_build_696, true, true)
}

func (dm_build_824 *dm_build_800) dm_build_787() error {
	return nil
}

func (dm_build_826 *dm_build_800) dm_build_788() {
	if dm_build_826.dm_build_804 == nil {
		dm_build_826.dm_build_801.dm_build_413.Dm_build_269(Dm_build_698, 0)
	} else {
		dm_build_826.dm_build_801.dm_build_413.Dm_build_269(Dm_build_698, dm_build_826.dm_build_804.id)
	}

	dm_build_826.dm_build_801.dm_build_413.Dm_build_265(Dm_build_699, dm_build_826.dm_build_802)
	dm_build_826.dm_build_801.dm_build_413.Dm_build_269(Dm_build_700, int32(dm_build_826.dm_build_801.dm_build_413.Dm_build_90()-Dm_build_696))
}

func (dm_build_828 *dm_build_800) dm_build_790() error {
	dm_build_828.dm_build_801.dm_build_413.Dm_build_95(0)
	dm_build_828.dm_build_801.dm_build_413.Dm_build_103(Dm_build_696, false, true)
	return dm_build_828.dm_build_833()
}

func (dm_build_830 *dm_build_800) dm_build_791() (interface{}, error) {
	return nil, nil
}

func (dm_build_832 *dm_build_800) dm_build_792() {
}

func (dm_build_834 *dm_build_800) dm_build_833() error {
	dm_build_834.dm_build_803 = dm_build_834.dm_build_801.dm_build_413.Dm_build_347(Dm_build_701)
	if dm_build_834.dm_build_803 < 0 && dm_build_834.dm_build_803 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_834.dm_build_803, dm_build_834.dm_build_835(), nil, ""}).throw()
	} else if dm_build_834.dm_build_803 > 0 {

	} else if dm_build_834.dm_build_802 == Dm_build_695 || dm_build_834.dm_build_802 == Dm_build_669 {
		dm_build_834.dm_build_835()
	}

	return nil
}

func (dm_build_836 *dm_build_800) dm_build_835() string {

	dm_build_837 := dm_build_836.dm_build_801.dm_build_414.getServerEncoding()

	if dm_build_837 != "" && dm_build_837 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_837 = ENCODING_GB18030
	}

	dm_build_836.dm_build_801.dm_build_413.Dm_build_103(int(dm_build_836.dm_build_801.dm_build_413.Dm_build_203()), false, true)

	dm_build_836.dm_build_801.dm_build_413.Dm_build_103(int(dm_build_836.dm_build_801.dm_build_413.Dm_build_203()), false, true)

	dm_build_836.dm_build_801.dm_build_413.Dm_build_103(int(dm_build_836.dm_build_801.dm_build_413.Dm_build_203()), false, true)

	return dm_build_836.dm_build_801.dm_build_413.Dm_build_245(dm_build_837, dm_build_836.dm_build_801.dm_build_414)
}

func (dm_build_839 *dm_build_800) dm_build_789(dm_build_840 dm_build_785) (dm_build_841 error) {
	dm_build_840.dm_build_786()
	if dm_build_841 = dm_build_840.dm_build_787(); dm_build_841 != nil {
		return dm_build_841
	}
	dm_build_840.dm_build_788()
	return nil
}

func (dm_build_843 *dm_build_800) dm_build_793(dm_build_844 dm_build_785) (dm_build_845 interface{}, dm_build_846 error) {
	dm_build_846 = dm_build_844.dm_build_790()
	if dm_build_846 != nil {
		return nil, dm_build_846
	}
	dm_build_845, dm_build_846 = dm_build_844.dm_build_791()
	if dm_build_846 != nil {
		return nil, dm_build_846
	}
	dm_build_844.dm_build_792()
	return dm_build_845, nil
}

func (dm_build_848 *dm_build_800) dm_build_794() {
	dm_build_848.dm_build_801.dm_build_413.Dm_build_261(Dm_build_704, dm_build_848.dm_build_796())
}

func (dm_build_850 *dm_build_800) dm_build_795() error {
	dm_build_851 := dm_build_850.dm_build_801.dm_build_413.Dm_build_341(Dm_build_704)
	dm_build_852 := dm_build_850.dm_build_796()
	if dm_build_851 != dm_build_852 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_854 *dm_build_800) dm_build_796() byte {
	dm_build_855 := dm_build_854.dm_build_801.dm_build_413.Dm_build_341(0)

	for i := 1; i < Dm_build_704; i++ {
		dm_build_855 ^= dm_build_854.dm_build_801.dm_build_413.Dm_build_341(i)
	}

	return dm_build_855
}

func (dm_build_857 *dm_build_800) dm_build_797() int32 {
	return dm_build_857.dm_build_801.dm_build_413.Dm_build_347(Dm_build_700)
}

func (dm_build_859 *dm_build_800) dm_build_798(dm_build_860 int32) {
	dm_build_859.dm_build_801.dm_build_413.Dm_build_269(Dm_build_700, dm_build_860)
}

func (dm_build_862 *dm_build_800) dm_build_799() int16 {
	return dm_build_862.dm_build_802
}

type dm_build_863 struct {
	dm_build_800
}

func dm_build_864(dm_build_865 *dm_build_410) *dm_build_863 {
	dm_build_866 := new(dm_build_863)
	dm_build_866.dm_build_805(dm_build_865, Dm_build_676)
	return dm_build_866
}

type dm_build_867 struct {
	dm_build_800
	dm_build_868 string
}

func dm_build_869(dm_build_870 *dm_build_410, dm_build_871 *DmStatement, dm_build_872 string) *dm_build_867 {
	dm_build_873 := new(dm_build_867)
	dm_build_873.dm_build_809(dm_build_870, Dm_build_684, dm_build_871)
	dm_build_873.dm_build_868 = dm_build_872
	dm_build_873.dm_build_804.cursorName = dm_build_872
	return dm_build_873
}

func (dm_build_875 *dm_build_867) dm_build_787() error {
	dm_build_875.dm_build_801.dm_build_413.Dm_build_191(dm_build_875.dm_build_868, dm_build_875.dm_build_801.dm_build_414.getServerEncoding(), dm_build_875.dm_build_801.dm_build_414)
	dm_build_875.dm_build_801.dm_build_413.Dm_build_129(1)
	return nil
}

type Dm_build_876 struct {
	dm_build_892
	dm_build_877 []OptParameter
}

func dm_build_878(dm_build_879 *dm_build_410, dm_build_880 *DmStatement, dm_build_881 []OptParameter) *Dm_build_876 {
	dm_build_882 := new(Dm_build_876)
	dm_build_882.dm_build_809(dm_build_879, Dm_build_694, dm_build_880)
	dm_build_882.dm_build_877 = dm_build_881
	return dm_build_882
}

func (dm_build_884 *Dm_build_876) dm_build_787() error {
	dm_build_885 := len(dm_build_884.dm_build_877)

	dm_build_884.dm_build_906(int16(dm_build_885), 1)

	dm_build_884.dm_build_801.dm_build_413.Dm_build_191(dm_build_884.dm_build_804.nativeSql, dm_build_884.dm_build_804.dmConn.getServerEncoding(), dm_build_884.dm_build_804.dmConn)

	for _, param := range dm_build_884.dm_build_877 {
		dm_build_884.dm_build_801.dm_build_413.Dm_build_121(param.ioType)
		dm_build_884.dm_build_801.dm_build_413.Dm_build_129(int32(param.tp))
		dm_build_884.dm_build_801.dm_build_413.Dm_build_129(int32(param.prec))
		dm_build_884.dm_build_801.dm_build_413.Dm_build_129(int32(param.scale))
	}

	for _, param := range dm_build_884.dm_build_877 {
		if param.bytes == nil {
			dm_build_884.dm_build_801.dm_build_413.Dm_build_137(Dm_build_712)
		} else {
			dm_build_884.dm_build_801.dm_build_413.Dm_build_167(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_887 *Dm_build_876) dm_build_791() (interface{}, error) {
	return dm_build_887.dm_build_892.dm_build_791()
}

const (
	Dm_build_888 int = 0x01

	Dm_build_889 int = 0x02

	Dm_build_890 int = 0x04

	Dm_build_891 int = 0x08
)

type dm_build_892 struct {
	dm_build_800
	dm_build_893 [][]interface{}
	dm_build_894 []parameter
	dm_build_895 bool
}

func dm_build_896(dm_build_897 *dm_build_410, dm_build_898 int16, dm_build_899 *DmStatement) *dm_build_892 {
	dm_build_900 := new(dm_build_892)
	dm_build_900.dm_build_809(dm_build_897, dm_build_898, dm_build_899)
	dm_build_900.dm_build_895 = true
	return dm_build_900
}

func dm_build_901(dm_build_902 *dm_build_410, dm_build_903 *DmStatement, dm_build_904 [][]interface{}) *dm_build_892 {
	dm_build_905 := new(dm_build_892)

	if dm_build_902.dm_build_414.Execute2 {
		dm_build_905.dm_build_809(dm_build_902, Dm_build_678, dm_build_903)
	} else {
		dm_build_905.dm_build_809(dm_build_902, Dm_build_674, dm_build_903)
	}

	dm_build_905.dm_build_894 = dm_build_903.params
	dm_build_905.dm_build_893 = dm_build_904
	dm_build_905.dm_build_895 = true
	return dm_build_905
}

func (dm_build_907 *dm_build_892) dm_build_906(dm_build_908 int16, dm_build_909 int64) {

	dm_build_910 := Dm_build_697

	if dm_build_907.dm_build_801.dm_build_414.autoCommit {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 1)
	} else {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 0)
	}

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_265(dm_build_910, dm_build_908)

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 1)

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_273(dm_build_910, dm_build_909)

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_273(dm_build_910, dm_build_907.dm_build_804.cursorUpdateRow)

	if dm_build_907.dm_build_804.maxRows <= 0 || dm_build_907.dm_build_804.dmConn.dmConnector.enRsCache {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_273(dm_build_910, INT64_MAX)
	} else {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_273(dm_build_910, dm_build_907.dm_build_804.maxRows)
	}

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 1)

	if dm_build_907.dm_build_801.dm_build_414.dmConnector.continueBatchOnError {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 1)
	} else {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 0)
	}

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 0)

	dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_261(dm_build_910, 0)

	if dm_build_907.dm_build_804.queryTimeout == 0 {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_269(dm_build_910, -1)
	} else {
		dm_build_910 += dm_build_907.dm_build_801.dm_build_413.Dm_build_269(dm_build_910, dm_build_907.dm_build_804.queryTimeout)
	}
}

func (dm_build_912 *dm_build_892) dm_build_787() error {
	var dm_build_913 int16
	var dm_build_914 int64

	if dm_build_912.dm_build_894 != nil {
		dm_build_913 = int16(len(dm_build_912.dm_build_894))
	} else {
		dm_build_913 = 0
	}

	if dm_build_912.dm_build_893 != nil {
		dm_build_914 = int64(len(dm_build_912.dm_build_893))
	} else {
		dm_build_914 = 0
	}

	dm_build_912.dm_build_906(dm_build_913, dm_build_914)

	if dm_build_913 > 0 {
		err := dm_build_912.dm_build_915(dm_build_912.dm_build_894)
		if err != nil {
			return err
		}
		if dm_build_912.dm_build_893 != nil && len(dm_build_912.dm_build_893) > 0 {
			for _, paramObject := range dm_build_912.dm_build_893 {
				if err := dm_build_912.dm_build_918(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_916 *dm_build_892) dm_build_915(dm_build_917 []parameter) error {
	for _, param := range dm_build_917 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_916.dm_build_801.dm_build_413.Dm_build_121(IO_TYPE_INOUT)
		} else {
			dm_build_916.dm_build_801.dm_build_413.Dm_build_121(param.ioType)
		}

		dm_build_916.dm_build_801.dm_build_413.Dm_build_129(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_916.dm_build_801.dm_build_413.Dm_build_129(lprec)

		dm_build_916.dm_build_801.dm_build_413.Dm_build_129(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_916.dm_build_801.dm_build_413)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_916.dm_build_801.dm_build_413)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_916.dm_build_801.dm_build_413)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_919 *dm_build_892) dm_build_918(dm_build_920 []interface{}) error {
	for i := 0; i < len(dm_build_919.dm_build_894); i++ {

		if dm_build_919.dm_build_894[i].colType == CURSOR {
			dm_build_919.dm_build_801.dm_build_413.Dm_build_125(ULINT_SIZE)
			dm_build_919.dm_build_801.dm_build_413.Dm_build_129(dm_build_919.dm_build_894[i].cursorStmt.id)
			continue
		}

		if dm_build_919.dm_build_894[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_920[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_920[i].([]byte); ok {
				if len(dataBytes) > Dm_build_707 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_919.dm_build_801.dm_build_413.Dm_build_167(dataBytes)
			}
		case int:
			if dm_build_920[i] == ParamDataEnum_Null {
				dm_build_919.dm_build_801.dm_build_413.Dm_build_137(Dm_build_712)
			} else if dm_build_920[i] == ParamDataEnum_OFF_ROW {
				dm_build_919.dm_build_801.dm_build_413.Dm_build_125(0)
			}
		case lobCtl:
			dm_build_919.dm_build_801.dm_build_413.Dm_build_137(uint16(Dm_build_711))
			dm_build_919.dm_build_801.dm_build_413.Dm_build_157(dm_build_920[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_922 *dm_build_892) dm_build_791() (interface{}, error) {
	dm_build_923 := execInfo{}
	dm_build_924 := dm_build_922.dm_build_804.dmConn

	dm_build_925 := Dm_build_697

	dm_build_923.retSqlType = dm_build_922.dm_build_801.dm_build_413.Dm_build_344(dm_build_925)
	dm_build_925 += USINT_SIZE

	dm_build_926 := dm_build_922.dm_build_801.dm_build_413.Dm_build_344(dm_build_925)
	dm_build_925 += USINT_SIZE

	dm_build_923.updateCount = dm_build_922.dm_build_801.dm_build_413.Dm_build_350(dm_build_925)
	dm_build_925 += DDWORD_SIZE

	dm_build_927 := dm_build_922.dm_build_801.dm_build_413.Dm_build_344(dm_build_925)
	dm_build_925 += USINT_SIZE

	dm_build_923.rsUpdatable = dm_build_922.dm_build_801.dm_build_413.Dm_build_341(dm_build_925) != 0
	dm_build_925 += BYTE_SIZE

	dm_build_928 := dm_build_922.dm_build_801.dm_build_413.Dm_build_344(dm_build_925)
	dm_build_925 += ULINT_SIZE

	dm_build_923.printLen = dm_build_922.dm_build_801.dm_build_413.Dm_build_347(dm_build_925)
	dm_build_925 += ULINT_SIZE

	var dm_build_929 int16 = -1
	if dm_build_923.retSqlType == Dm_build_762 || dm_build_923.retSqlType == Dm_build_763 {
		dm_build_923.rowid = 0

		dm_build_923.rsBdta = dm_build_922.dm_build_801.dm_build_413.Dm_build_341(dm_build_925) == Dm_build_775
		dm_build_925 += BYTE_SIZE

		dm_build_929 = dm_build_922.dm_build_801.dm_build_413.Dm_build_344(dm_build_925)
		dm_build_925 += USINT_SIZE
		dm_build_925 += 5
	} else {
		dm_build_923.rowid = dm_build_922.dm_build_801.dm_build_413.Dm_build_350(dm_build_925)
		dm_build_925 += DDWORD_SIZE
	}

	dm_build_923.execId = dm_build_922.dm_build_801.dm_build_413.Dm_build_347(dm_build_925)
	dm_build_925 += ULINT_SIZE

	dm_build_923.rsCacheOffset = dm_build_922.dm_build_801.dm_build_413.Dm_build_347(dm_build_925)
	dm_build_925 += ULINT_SIZE

	dm_build_930 := dm_build_922.dm_build_801.dm_build_413.Dm_build_341(dm_build_925)
	dm_build_925 += BYTE_SIZE
	dm_build_931 := (dm_build_930 & 0x01) == 0x01
	dm_build_932 := (dm_build_930 & 0x02) == 0x02

	dm_build_924.TrxStatus = dm_build_922.dm_build_801.dm_build_413.Dm_build_347(dm_build_925)
	dm_build_924.setTrxFinish(dm_build_924.TrxStatus)
	dm_build_925 += ULINT_SIZE

	if dm_build_923.printLen > 0 {
		bytes := dm_build_922.dm_build_801.dm_build_413.Dm_build_224(int(dm_build_923.printLen))
		dm_build_923.printMsg = Dm_build_1261.Dm_build_1415(bytes, 0, len(bytes), dm_build_924.getServerEncoding(), dm_build_924)
	}

	if dm_build_927 > 0 {
		dm_build_923.outParamDatas = dm_build_922.dm_build_933(int(dm_build_927))
	}

	switch dm_build_923.retSqlType {
	case Dm_build_764:
		dm_build_924.dmConnector.localTimezone = dm_build_922.dm_build_801.dm_build_413.Dm_build_200()
	case Dm_build_762:
		dm_build_923.hasResultSet = true
		if dm_build_926 > 0 {
			dm_build_922.dm_build_804.columns = dm_build_922.dm_build_942(int(dm_build_926), dm_build_923.rsBdta)
		}
		dm_build_922.dm_build_952(&dm_build_923, len(dm_build_922.dm_build_804.columns), int(dm_build_928), int(dm_build_929))
	case Dm_build_763:
		if dm_build_926 > 0 || dm_build_928 > 0 {
			dm_build_923.hasResultSet = true
		}
		if dm_build_926 > 0 {
			dm_build_922.dm_build_804.columns = dm_build_922.dm_build_942(int(dm_build_926), dm_build_923.rsBdta)
		}
		dm_build_922.dm_build_952(&dm_build_923, len(dm_build_922.dm_build_804.columns), int(dm_build_928), int(dm_build_929))
	case Dm_build_765:
		dm_build_924.IsoLevel = int32(dm_build_922.dm_build_801.dm_build_413.Dm_build_200())
		dm_build_924.ReadOnly = dm_build_922.dm_build_801.dm_build_413.Dm_build_197() == 1
	case Dm_build_758:
		dm_build_924.Schema = dm_build_922.dm_build_801.dm_build_413.Dm_build_245(dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_755:
		dm_build_923.explain = dm_build_922.dm_build_801.dm_build_413.Dm_build_245(dm_build_924.getServerEncoding(), dm_build_924)

	case Dm_build_759, Dm_build_761, Dm_build_760:
		if dm_build_931 {

			counts := dm_build_922.dm_build_801.dm_build_413.Dm_build_203()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_922.dm_build_801.dm_build_413.Dm_build_206()
			}
			dm_build_923.updateCounts = rowCounts
		}

		if dm_build_932 {
			rows := dm_build_922.dm_build_801.dm_build_413.Dm_build_203()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_922.dm_build_801.dm_build_413.Dm_build_206()
			}
			dm_build_923.lastInsertId = lastInsertId

		} else if dm_build_923.updateCount == 1 {
			dm_build_923.lastInsertId = dm_build_923.rowid
		}

		if dm_build_922.dm_build_803 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_922.dm_build_958(dm_build_923.updateCounts)
		}
	case Dm_build_768:
		len := dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
		dm_build_924.OracleDateFormat = dm_build_922.dm_build_801.dm_build_413.Dm_build_240(int(len), dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_770:

		len := dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
		dm_build_924.OracleTimestampFormat = dm_build_922.dm_build_801.dm_build_413.Dm_build_240(int(len), dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_771:

		len := dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
		dm_build_924.OracleTimestampTZFormat = dm_build_922.dm_build_801.dm_build_413.Dm_build_240(int(len), dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_769:
		len := dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
		dm_build_924.OracleTimeFormat = dm_build_922.dm_build_801.dm_build_413.Dm_build_240(int(len), dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_772:
		len := dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
		dm_build_924.OracleTimeTZFormat = dm_build_922.dm_build_801.dm_build_413.Dm_build_240(int(len), dm_build_924.getServerEncoding(), dm_build_924)
	case Dm_build_773:
		dm_build_924.OracleDateLanguage = dm_build_922.dm_build_801.dm_build_413.Dm_build_215()
	}

	return &dm_build_923, nil
}

func (dm_build_934 *dm_build_892) dm_build_933(dm_build_935 int) [][]byte {
	dm_build_936 := make([]int, dm_build_935)

	dm_build_937 := 0
	for i := 0; i < len(dm_build_934.dm_build_894); i++ {
		if dm_build_934.dm_build_894[i].ioType == IO_TYPE_INOUT || dm_build_934.dm_build_894[i].ioType == IO_TYPE_OUT {
			dm_build_936[dm_build_937] = i
			dm_build_937++
		}
	}

	dm_build_938 := make([][]byte, len(dm_build_934.dm_build_894))
	var dm_build_939 int32
	var dm_build_940 bool
	var dm_build_941 []byte = nil
	for i := 0; i < dm_build_935; i++ {
		dm_build_940 = false
		dm_build_939 = int32(dm_build_934.dm_build_801.dm_build_413.Dm_build_218())

		if dm_build_939 == int32(Dm_build_712) {
			dm_build_939 = 0
			dm_build_940 = true
		} else if dm_build_939 == int32(Dm_build_713) {
			dm_build_939 = dm_build_934.dm_build_801.dm_build_413.Dm_build_203()
		}

		if dm_build_940 {
			dm_build_938[dm_build_936[i]] = nil
		} else {
			dm_build_941 = dm_build_934.dm_build_801.dm_build_413.Dm_build_224(int(dm_build_939))
			dm_build_938[dm_build_936[i]] = dm_build_941
		}
	}

	return dm_build_938
}

func (dm_build_943 *dm_build_892) dm_build_942(dm_build_944 int, dm_build_945 bool) []column {
	dm_build_946 := dm_build_943.dm_build_801.dm_build_414.getServerEncoding()
	var dm_build_947, dm_build_948, dm_build_949, dm_build_950 int16
	dm_build_951 := make([]column, dm_build_944)
	for i := 0; i < dm_build_944; i++ {
		dm_build_951[i].InitColumn()

		dm_build_951[i].colType = dm_build_943.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_951[i].prec = dm_build_943.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_951[i].scale = dm_build_943.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_951[i].nullable = dm_build_943.dm_build_801.dm_build_413.Dm_build_203() != 0

		itemFlag := dm_build_943.dm_build_801.dm_build_413.Dm_build_200()
		dm_build_951[i].lob = int(itemFlag)&Dm_build_889 != 0
		dm_build_951[i].identity = int(itemFlag)&Dm_build_888 != 0
		dm_build_951[i].readonly = int(itemFlag)&Dm_build_890 != 0

		dm_build_943.dm_build_801.dm_build_413.Dm_build_103(4, false, true)

		dm_build_943.dm_build_801.dm_build_413.Dm_build_103(2, false, true)

		dm_build_947 = dm_build_943.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_948 = dm_build_943.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_949 = dm_build_943.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_950 = dm_build_943.dm_build_801.dm_build_413.Dm_build_200()
		dm_build_951[i].name = dm_build_943.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_947), dm_build_946, dm_build_943.dm_build_801.dm_build_414)
		dm_build_951[i].typeName = dm_build_943.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_948), dm_build_946, dm_build_943.dm_build_801.dm_build_414)
		dm_build_951[i].tableName = dm_build_943.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_949), dm_build_946, dm_build_943.dm_build_801.dm_build_414)
		dm_build_951[i].schemaName = dm_build_943.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_950), dm_build_946, dm_build_943.dm_build_801.dm_build_414)

		if dm_build_943.dm_build_804.readBaseColName {
			dm_build_951[i].baseName = dm_build_943.dm_build_801.dm_build_413.Dm_build_253(dm_build_946, dm_build_943.dm_build_801.dm_build_414)
		}

		if dm_build_951[i].lob {
			dm_build_951[i].lobTabId = dm_build_943.dm_build_801.dm_build_413.Dm_build_203()
			dm_build_951[i].lobColId = dm_build_943.dm_build_801.dm_build_413.Dm_build_200()
		}

	}

	for i := 0; i < dm_build_944; i++ {

		if isComplexType(int(dm_build_951[i].colType), int(dm_build_951[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_943.dm_build_801.dm_build_414)
			strDesc.unpack(dm_build_943.dm_build_801.dm_build_413)
			dm_build_951[i].typeDescriptor = strDesc
		}
	}

	return dm_build_951
}

func (dm_build_953 *dm_build_892) dm_build_952(dm_build_954 *execInfo, dm_build_955 int, dm_build_956 int, dm_build_957 int) {
	if dm_build_956 > 0 {
		startOffset := dm_build_953.dm_build_801.dm_build_413.Dm_build_98()
		if dm_build_954.rsBdta {
			dm_build_954.rsDatas = dm_build_953.dm_build_971(dm_build_953.dm_build_804.columns, dm_build_957)
		} else {
			datas := make([][][]byte, dm_build_956)

			for i := 0; i < dm_build_956; i++ {

				datas[i] = make([][]byte, dm_build_955+1)

				dm_build_953.dm_build_801.dm_build_413.Dm_build_103(2, false, true)

				datas[i][0] = dm_build_953.dm_build_801.dm_build_413.Dm_build_224(LINT64_SIZE)

				dm_build_953.dm_build_801.dm_build_413.Dm_build_103(2*dm_build_955, false, true)

				for j := 1; j < dm_build_955+1; j++ {

					colLen := dm_build_953.dm_build_801.dm_build_413.Dm_build_218()
					if colLen == Dm_build_716 {
						datas[i][j] = nil
					} else if colLen != Dm_build_717 {
						datas[i][j] = dm_build_953.dm_build_801.dm_build_413.Dm_build_224(int(colLen))
					} else {
						datas[i][j] = dm_build_953.dm_build_801.dm_build_413.Dm_build_228()
					}
				}
			}

			dm_build_954.rsDatas = datas
		}
		dm_build_954.rsSizeof = dm_build_953.dm_build_801.dm_build_413.Dm_build_98() - startOffset
	}

	if dm_build_954.rsCacheOffset > 0 {
		tbCount := dm_build_953.dm_build_801.dm_build_413.Dm_build_200()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_953.dm_build_801.dm_build_413.Dm_build_203()
			tss[i] = dm_build_953.dm_build_801.dm_build_413.Dm_build_206()
		}

		dm_build_954.tbIds = ids[:]
		dm_build_954.tbTss = tss[:]
	}
}

func (dm_build_959 *dm_build_892) dm_build_958(dm_build_960 []int64) error {

	dm_build_959.dm_build_801.dm_build_413.Dm_build_103(4, false, true)

	dm_build_961 := dm_build_959.dm_build_801.dm_build_413.Dm_build_203()

	dm_build_962 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_961); i++ {
		irow := dm_build_959.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_960[irow] = -3

		code := dm_build_959.dm_build_801.dm_build_413.Dm_build_203()

		errStr := dm_build_959.dm_build_801.dm_build_413.Dm_build_253(dm_build_959.dm_build_801.dm_build_414.getServerEncoding(), dm_build_959.dm_build_801.dm_build_414)

		dm_build_962 = append(dm_build_962, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_962) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_962 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_963 = 0

	Dm_build_964 = Dm_build_963 + ULINT_SIZE

	Dm_build_965 = Dm_build_964 + USINT_SIZE

	Dm_build_966 = Dm_build_965 + ULINT_SIZE

	Dm_build_967 = Dm_build_966 + ULINT_SIZE

	Dm_build_968 = Dm_build_967 + BYTE_SIZE

	Dm_build_969 = -2

	Dm_build_970 = -3
)

func (dm_build_972 *dm_build_892) dm_build_971(dm_build_973 []column, dm_build_974 int) [][][]byte {

	dm_build_975 := dm_build_972.dm_build_801.dm_build_413.Dm_build_221()

	dm_build_976 := dm_build_972.dm_build_801.dm_build_413.Dm_build_218()

	var dm_build_977 bool

	if dm_build_974 >= 0 && int(dm_build_976) == len(dm_build_973)+1 {
		dm_build_977 = true
	} else {
		dm_build_977 = false
	}

	dm_build_972.dm_build_801.dm_build_413.Dm_build_103(ULINT_SIZE, false, true)

	dm_build_972.dm_build_801.dm_build_413.Dm_build_103(ULINT_SIZE, false, true)

	dm_build_972.dm_build_801.dm_build_413.Dm_build_103(BYTE_SIZE, false, true)

	dm_build_978 := make([]uint16, dm_build_976)
	for icol := 0; icol < int(dm_build_976); icol++ {
		dm_build_978[icol] = dm_build_972.dm_build_801.dm_build_413.Dm_build_218()
	}

	dm_build_979 := make([]uint32, dm_build_976)
	dm_build_980 := make([][][]byte, dm_build_975)

	for i := uint32(0); i < dm_build_975; i++ {
		dm_build_980[i] = make([][]byte, len(dm_build_973)+1)
	}

	for icol := 0; icol < int(dm_build_976); icol++ {
		dm_build_979[icol] = dm_build_972.dm_build_801.dm_build_413.Dm_build_221()
	}

	for icol := 0; icol < int(dm_build_976); icol++ {

		dataCol := icol + 1
		if dm_build_977 && icol == dm_build_974 {
			dataCol = 0
		} else if dm_build_977 && icol > dm_build_974 {
			dataCol = icol
		}

		allNotNull := dm_build_972.dm_build_801.dm_build_413.Dm_build_203() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_975)
			for irow := uint32(0); irow < dm_build_975; irow++ {
				isNull[irow] = dm_build_972.dm_build_801.dm_build_413.Dm_build_197() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_975; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_980[irow][dataCol] = dm_build_972.dm_build_981(int(dm_build_978[icol]))
			}
		}
	}

	if !dm_build_977 && dm_build_974 >= 0 {
		for irow := uint32(0); irow < dm_build_975; irow++ {
			dm_build_980[irow][0] = dm_build_980[irow][dm_build_974+1]
		}
	}

	return dm_build_980
}

func (dm_build_982 *dm_build_892) dm_build_981(dm_build_983 int) []byte {

	dm_build_984 := dm_build_982.dm_build_987(dm_build_983)

	dm_build_985 := int32(0)
	if dm_build_984 == Dm_build_969 {
		dm_build_985 = dm_build_982.dm_build_801.dm_build_413.Dm_build_203()
		dm_build_984 = int(dm_build_982.dm_build_801.dm_build_413.Dm_build_203())
	} else if dm_build_984 == Dm_build_970 {
		dm_build_984 = int(dm_build_982.dm_build_801.dm_build_413.Dm_build_203())
	}

	dm_build_986 := dm_build_982.dm_build_801.dm_build_413.Dm_build_224(dm_build_984 + int(dm_build_985))
	if dm_build_985 == 0 {
		return dm_build_986
	}

	for i := dm_build_984; i < len(dm_build_986); i++ {
		dm_build_986[i] = ' '
	}
	return dm_build_986
}

func (dm_build_988 *dm_build_892) dm_build_987(dm_build_989 int) int {

	dm_build_990 := 0
	switch dm_build_989 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_990 = 4

	case BIGINT:

		dm_build_990 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_990 = Dm_build_969

	case DECIMAL:
		dm_build_990 = Dm_build_970

	case REAL:
		dm_build_990 = 4

	case DOUBLE:
		dm_build_990 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_990 = 12

	case INTERVAL_YM:
		dm_build_990 = 12

	case INTERVAL_DT:
		dm_build_990 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_990
}

const (
	Dm_build_991 = Dm_build_697

	Dm_build_992 = Dm_build_991 + DDWORD_SIZE

	Dm_build_993 = Dm_build_992 + LINT64_SIZE

	Dm_build_994 = Dm_build_993 + USINT_SIZE

	Dm_build_995 = Dm_build_697

	Dm_build_996 = Dm_build_995 + DDWORD_SIZE
)

type dm_build_997 struct {
	dm_build_892
	dm_build_998  *innerRows
	dm_build_999  int64
	dm_build_1000 int64
}

func dm_build_1001(dm_build_1002 *dm_build_410, dm_build_1003 *innerRows, dm_build_1004 int64, dm_build_1005 int64) *dm_build_997 {
	dm_build_1006 := new(dm_build_997)
	dm_build_1006.dm_build_809(dm_build_1002, Dm_build_675, dm_build_1003.dmStmt)
	dm_build_1006.dm_build_998 = dm_build_1003
	dm_build_1006.dm_build_999 = dm_build_1004
	dm_build_1006.dm_build_1000 = dm_build_1005
	return dm_build_1006
}

func (dm_build_1008 *dm_build_997) dm_build_787() error {

	dm_build_1008.dm_build_801.dm_build_413.Dm_build_273(Dm_build_991, dm_build_1008.dm_build_999)

	dm_build_1008.dm_build_801.dm_build_413.Dm_build_273(Dm_build_992, dm_build_1008.dm_build_1000)

	dm_build_1008.dm_build_801.dm_build_413.Dm_build_265(Dm_build_993, dm_build_1008.dm_build_998.id)

	dm_build_1009 := dm_build_1008.dm_build_998.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_1010 int32
	if dm_build_1008.dm_build_998.sizeOfRow != 0 && dm_build_1008.dm_build_998.fetchSize != 0 {
		if dm_build_1008.dm_build_998.sizeOfRow*dm_build_1008.dm_build_998.fetchSize > int(INT32_MAX) {
			dm_build_1010 = INT32_MAX
		} else {
			dm_build_1010 = int32(dm_build_1008.dm_build_998.sizeOfRow * dm_build_1008.dm_build_998.fetchSize)
		}

		if dm_build_1010 < Dm_build_728 {
			dm_build_1009 = int(Dm_build_728)
		} else if dm_build_1010 > Dm_build_729 {
			dm_build_1009 = int(Dm_build_729)
		} else {
			dm_build_1009 = int(dm_build_1010)
		}

		dm_build_1008.dm_build_801.dm_build_413.Dm_build_269(Dm_build_994, int32(dm_build_1009))
	}

	return nil
}

func (dm_build_1012 *dm_build_997) dm_build_791() (interface{}, error) {
	dm_build_1013 := execInfo{}
	dm_build_1013.rsBdta = dm_build_1012.dm_build_998.isBdta

	dm_build_1013.updateCount = dm_build_1012.dm_build_801.dm_build_413.Dm_build_350(Dm_build_995)
	dm_build_1014 := dm_build_1012.dm_build_801.dm_build_413.Dm_build_347(Dm_build_996)

	dm_build_1012.dm_build_952(&dm_build_1013, len(dm_build_1012.dm_build_998.columns), int(dm_build_1014), -1)

	return &dm_build_1013, nil
}

type dm_build_1015 struct {
	dm_build_800
	dm_build_1016 *lob
	dm_build_1017 int
	dm_build_1018 int
}

func dm_build_1019(dm_build_1020 *dm_build_410, dm_build_1021 *lob, dm_build_1022 int, dm_build_1023 int) *dm_build_1015 {
	dm_build_1024 := new(dm_build_1015)
	dm_build_1024.dm_build_805(dm_build_1020, Dm_build_688)
	dm_build_1024.dm_build_1016 = dm_build_1021
	dm_build_1024.dm_build_1017 = dm_build_1022
	dm_build_1024.dm_build_1018 = dm_build_1023
	return dm_build_1024
}

func (dm_build_1026 *dm_build_1015) dm_build_787() error {

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_121(byte(dm_build_1026.dm_build_1016.lobFlag))

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(dm_build_1026.dm_build_1016.tabId)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.colId)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1026.dm_build_1016.blobId))

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.groupId)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.fileId)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(dm_build_1026.dm_build_1016.pageNo)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.curFileId)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(dm_build_1026.dm_build_1016.curPageNo)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(dm_build_1026.dm_build_1016.totalOffset)

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(int32(dm_build_1026.dm_build_1017))

	dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(int32(dm_build_1026.dm_build_1018))

	if dm_build_1026.dm_build_801.dm_build_414.NewLobFlag {
		dm_build_1026.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1026.dm_build_1016.rowId))
		dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.exGroupId)
		dm_build_1026.dm_build_801.dm_build_413.Dm_build_125(dm_build_1026.dm_build_1016.exFileId)
		dm_build_1026.dm_build_801.dm_build_413.Dm_build_129(dm_build_1026.dm_build_1016.exPageNo)
	}

	return nil
}

func (dm_build_1028 *dm_build_1015) dm_build_791() (interface{}, error) {

	dm_build_1028.dm_build_1016.readOver = dm_build_1028.dm_build_801.dm_build_413.Dm_build_197() == 1
	var dm_build_1029 = dm_build_1028.dm_build_801.dm_build_413.Dm_build_221()
	if dm_build_1029 <= 0 {
		return []byte(nil), nil
	}
	dm_build_1028.dm_build_1016.curFileId = dm_build_1028.dm_build_801.dm_build_413.Dm_build_200()
	dm_build_1028.dm_build_1016.curPageNo = dm_build_1028.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1028.dm_build_1016.totalOffset = dm_build_1028.dm_build_801.dm_build_413.Dm_build_203()

	return dm_build_1028.dm_build_801.dm_build_413.Dm_build_234(int(dm_build_1029)), nil
}

type dm_build_1030 struct {
	dm_build_800
	dm_build_1031 *lob
}

func dm_build_1032(dm_build_1033 *dm_build_410, dm_build_1034 *lob) *dm_build_1030 {
	dm_build_1035 := new(dm_build_1030)
	dm_build_1035.dm_build_805(dm_build_1033, Dm_build_685)
	dm_build_1035.dm_build_1031 = dm_build_1034
	return dm_build_1035
}

func (dm_build_1037 *dm_build_1030) dm_build_787() error {

	dm_build_1037.dm_build_801.dm_build_413.Dm_build_121(byte(dm_build_1037.dm_build_1031.lobFlag))

	dm_build_1037.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1037.dm_build_1031.blobId))

	dm_build_1037.dm_build_801.dm_build_413.Dm_build_125(dm_build_1037.dm_build_1031.groupId)

	dm_build_1037.dm_build_801.dm_build_413.Dm_build_125(dm_build_1037.dm_build_1031.fileId)

	dm_build_1037.dm_build_801.dm_build_413.Dm_build_129(dm_build_1037.dm_build_1031.pageNo)

	if dm_build_1037.dm_build_801.dm_build_414.NewLobFlag {
		dm_build_1037.dm_build_801.dm_build_413.Dm_build_129(dm_build_1037.dm_build_1031.tabId)
		dm_build_1037.dm_build_801.dm_build_413.Dm_build_125(dm_build_1037.dm_build_1031.colId)
		dm_build_1037.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1037.dm_build_1031.rowId))

		dm_build_1037.dm_build_801.dm_build_413.Dm_build_125(dm_build_1037.dm_build_1031.exGroupId)
		dm_build_1037.dm_build_801.dm_build_413.Dm_build_125(dm_build_1037.dm_build_1031.exFileId)
		dm_build_1037.dm_build_801.dm_build_413.Dm_build_129(dm_build_1037.dm_build_1031.exPageNo)
	}

	return nil
}

func (dm_build_1039 *dm_build_1030) dm_build_791() (interface{}, error) {

	return int64(dm_build_1039.dm_build_801.dm_build_413.Dm_build_203()), nil
}

type dm_build_1040 struct {
	dm_build_800
	dm_build_1041 *lob
	dm_build_1042 int
}

func dm_build_1043(dm_build_1044 *dm_build_410, dm_build_1045 *lob, dm_build_1046 int) *dm_build_1040 {
	dm_build_1047 := new(dm_build_1040)
	dm_build_1047.dm_build_805(dm_build_1044, Dm_build_687)
	dm_build_1047.dm_build_1041 = dm_build_1045
	dm_build_1047.dm_build_1042 = dm_build_1046
	return dm_build_1047
}

func (dm_build_1049 *dm_build_1040) dm_build_787() error {

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_121(byte(dm_build_1049.dm_build_1041.lobFlag))

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1049.dm_build_1041.blobId))

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_125(dm_build_1049.dm_build_1041.groupId)

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_125(dm_build_1049.dm_build_1041.fileId)

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_129(dm_build_1049.dm_build_1041.pageNo)

	dm_build_1049.dm_build_801.dm_build_413.Dm_build_129(dm_build_1049.dm_build_1041.tabId)
	dm_build_1049.dm_build_801.dm_build_413.Dm_build_125(dm_build_1049.dm_build_1041.colId)
	dm_build_1049.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1049.dm_build_1041.rowId))
	dm_build_1049.dm_build_801.dm_build_413.Dm_build_157(Dm_build_1261.Dm_build_1460(uint32(dm_build_1049.dm_build_1042)))

	if dm_build_1049.dm_build_801.dm_build_414.NewLobFlag {
		dm_build_1049.dm_build_801.dm_build_413.Dm_build_125(dm_build_1049.dm_build_1041.exGroupId)
		dm_build_1049.dm_build_801.dm_build_413.Dm_build_125(dm_build_1049.dm_build_1041.exFileId)
		dm_build_1049.dm_build_801.dm_build_413.Dm_build_129(dm_build_1049.dm_build_1041.exPageNo)
	}
	return nil
}

func (dm_build_1051 *dm_build_1040) dm_build_791() (interface{}, error) {

	dm_build_1052 := dm_build_1051.dm_build_801.dm_build_413.Dm_build_221()
	dm_build_1051.dm_build_1041.blobId = dm_build_1051.dm_build_801.dm_build_413.Dm_build_206()
	dm_build_1051.dm_build_1041.resetCurrentInfo()
	return int64(dm_build_1052), nil
}

const (
	Dm_build_1053 = Dm_build_697

	Dm_build_1054 = Dm_build_1053 + ULINT_SIZE

	Dm_build_1055 = Dm_build_1054 + ULINT_SIZE

	Dm_build_1056 = Dm_build_1055 + ULINT_SIZE

	Dm_build_1057 = Dm_build_1056 + BYTE_SIZE

	Dm_build_1058 = Dm_build_1057 + USINT_SIZE

	Dm_build_1059 = Dm_build_1058 + ULINT_SIZE

	Dm_build_1060 = Dm_build_1059 + BYTE_SIZE

	Dm_build_1061 = Dm_build_1060 + BYTE_SIZE

	Dm_build_1062 = Dm_build_1061 + BYTE_SIZE

	Dm_build_1063 = Dm_build_697

	Dm_build_1064 = Dm_build_1063 + ULINT_SIZE

	Dm_build_1065 = Dm_build_1064 + ULINT_SIZE

	Dm_build_1066 = Dm_build_1065 + BYTE_SIZE

	Dm_build_1067 = Dm_build_1066 + ULINT_SIZE

	Dm_build_1068 = Dm_build_1067 + BYTE_SIZE

	Dm_build_1069 = Dm_build_1068 + BYTE_SIZE

	Dm_build_1070 = Dm_build_1069 + USINT_SIZE

	Dm_build_1071 = Dm_build_1070 + USINT_SIZE

	Dm_build_1072 = Dm_build_1071 + BYTE_SIZE

	Dm_build_1073 = Dm_build_1072 + USINT_SIZE

	Dm_build_1074 = Dm_build_1073 + BYTE_SIZE

	Dm_build_1075 = Dm_build_1074 + BYTE_SIZE

	Dm_build_1076 = Dm_build_1075 + ULINT_SIZE
)

type dm_build_1077 struct {
	dm_build_800

	dm_build_1078 *DmConnection

	dm_build_1079 bool
}

func dm_build_1080(dm_build_1081 *dm_build_410) *dm_build_1077 {
	dm_build_1082 := new(dm_build_1077)
	dm_build_1082.dm_build_805(dm_build_1081, Dm_build_669)
	dm_build_1082.dm_build_1078 = dm_build_1081.dm_build_414
	return dm_build_1082
}

func (dm_build_1084 *dm_build_1077) dm_build_787() error {

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1053, Dm_build_708)

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1054, g2dbIsoLevel(dm_build_1084.dm_build_1078.IsoLevel))
	dm_build_1084.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1055, int32(Locale))
	dm_build_1084.dm_build_801.dm_build_413.Dm_build_265(Dm_build_1057, dm_build_1084.dm_build_1078.dmConnector.localTimezone)

	if dm_build_1084.dm_build_1078.ReadOnly {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1056, Dm_build_731)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1056, Dm_build_730)
	}

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1058, int32(dm_build_1084.dm_build_1078.dmConnector.sessionTimeout))

	if dm_build_1084.dm_build_1078.dmConnector.mppLocal {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1059, 1)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1059, 0)
	}

	if dm_build_1084.dm_build_1078.dmConnector.rwSeparate {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1060, 1)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1060, 0)
	}

	if dm_build_1084.dm_build_1078.NewLobFlag {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1061, 1)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1061, 0)
	}

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1062, dm_build_1084.dm_build_1078.dmConnector.osAuthType)

	dm_build_1085 := dm_build_1084.dm_build_1078.getServerEncoding()

	if dm_build_1084.dm_build_801.dm_build_419 != "" {

	}

	dm_build_1086 := Dm_build_1261.Dm_build_1471(dm_build_1084.dm_build_1078.dmConnector.user, dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)
	dm_build_1087 := Dm_build_1261.Dm_build_1471(dm_build_1084.dm_build_1078.dmConnector.password, dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)
	if len(dm_build_1086) > Dm_build_705 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_1087) > Dm_build_705 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_1084.dm_build_801.dm_build_416 && dm_build_1084.dm_build_1078.dmConnector.loginCertificate != "" {

	} else if dm_build_1084.dm_build_801.dm_build_416 {
		dm_build_1086 = dm_build_1084.dm_build_801.dm_build_415.Encrypt(dm_build_1086, false)
		dm_build_1087 = dm_build_1084.dm_build_801.dm_build_415.Encrypt(dm_build_1087, false)
	}

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_161(dm_build_1086)
	dm_build_1084.dm_build_801.dm_build_413.Dm_build_161(dm_build_1087)

	dm_build_1084.dm_build_801.dm_build_413.Dm_build_173(dm_build_1084.dm_build_1078.dmConnector.appName, dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)
	dm_build_1084.dm_build_801.dm_build_413.Dm_build_173(dm_build_1084.dm_build_1078.dmConnector.osName, dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_173(hostName, dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_173("", dm_build_1085, dm_build_1084.dm_build_801.dm_build_414)
	}

	if dm_build_1084.dm_build_1078.dmConnector.rwStandby {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_121(1)
	} else {
		dm_build_1084.dm_build_801.dm_build_413.Dm_build_121(0)
	}

	return nil
}

func (dm_build_1089 *dm_build_1077) dm_build_791() (interface{}, error) {

	dm_build_1089.dm_build_1078.MaxRowSize = dm_build_1089.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1063)
	dm_build_1089.dm_build_1078.DDLAutoCommit = dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1065) == 1
	dm_build_1089.dm_build_1078.IsoLevel = dm_build_1089.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1066)
	dm_build_1089.dm_build_1078.dmConnector.caseSensitive = dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1067) == 1
	dm_build_1089.dm_build_1078.BackslashEscape = dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1068) == 1
	dm_build_1089.dm_build_1078.SvrStat = dm_build_1089.dm_build_801.dm_build_413.Dm_build_344(Dm_build_1070)
	dm_build_1089.dm_build_1078.SvrMode = dm_build_1089.dm_build_801.dm_build_413.Dm_build_344(Dm_build_1069)
	dm_build_1089.dm_build_1078.ConstParaOpt = dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1071) == 1
	dm_build_1089.dm_build_1078.DbTimezone = dm_build_1089.dm_build_801.dm_build_413.Dm_build_344(Dm_build_1072)
	dm_build_1089.dm_build_1078.NewLobFlag = dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1074) == 1

	if dm_build_1089.dm_build_1078.dmConnector.bufPrefetch == 0 {
		dm_build_1089.dm_build_1078.dmConnector.bufPrefetch = int(dm_build_1089.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1075))
	}

	dm_build_1089.dm_build_1078.LifeTimeRemainder = dm_build_1089.dm_build_801.dm_build_413.Dm_build_344(Dm_build_1076)

	dm_build_1090 := dm_build_1089.dm_build_1078.getServerEncoding()

	dm_build_1089.dm_build_1078.InstanceName = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
	dm_build_1089.dm_build_1078.Schema = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
	dm_build_1089.dm_build_1078.LastLoginIP = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
	dm_build_1089.dm_build_1078.LastLoginTime = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
	dm_build_1089.dm_build_1078.FailedAttempts = dm_build_1089.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1089.dm_build_1078.LoginWarningID = dm_build_1089.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1089.dm_build_1078.GraceTimeRemainder = dm_build_1089.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1089.dm_build_1078.Guid = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
	dm_build_1089.dm_build_1078.DbName = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)

	if dm_build_1089.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1073) == 1 {
		dm_build_1089.dm_build_1078.StandbyHost = dm_build_1089.dm_build_801.dm_build_413.Dm_build_245(dm_build_1090, dm_build_1089.dm_build_801.dm_build_414)
		dm_build_1089.dm_build_1078.StandbyPort = dm_build_1089.dm_build_801.dm_build_413.Dm_build_203()
		dm_build_1089.dm_build_1078.StandbyCount = dm_build_1089.dm_build_801.dm_build_413.Dm_build_218()
	}

	if dm_build_1089.dm_build_801.dm_build_413.Dm_build_100(false) > 0 {
		dm_build_1089.dm_build_1078.SessionID = dm_build_1089.dm_build_801.dm_build_413.Dm_build_206()
	}

	if dm_build_1089.dm_build_801.dm_build_413.Dm_build_100(false) > 0 {
		if dm_build_1089.dm_build_801.dm_build_413.Dm_build_197() == 1 {

			dm_build_1089.dm_build_1078.OracleDateFormat = "DD-MON-YY"

			dm_build_1089.dm_build_1078.OracleTimestampFormat = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_1089.dm_build_1078.OracleTimestampTZFormat = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_1089.dm_build_1078.OracleTimeFormat = "HH12.MI.SS.FF6 AM"

			dm_build_1089.dm_build_1078.OracleTimeTZFormat = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	return nil, nil
}

const (
	Dm_build_1091 = Dm_build_697
)

type dm_build_1092 struct {
	dm_build_892
	dm_build_1093 int16
}

func dm_build_1094(dm_build_1095 *dm_build_410, dm_build_1096 *DmStatement, dm_build_1097 int16) *dm_build_1092 {
	dm_build_1098 := new(dm_build_1092)
	dm_build_1098.dm_build_809(dm_build_1095, Dm_build_689, dm_build_1096)
	dm_build_1098.dm_build_1093 = dm_build_1097
	return dm_build_1098
}

func (dm_build_1100 *dm_build_1092) dm_build_787() error {
	dm_build_1100.dm_build_801.dm_build_413.Dm_build_265(Dm_build_1091, dm_build_1100.dm_build_1093)
	return nil
}

func (dm_build_1102 *dm_build_1092) dm_build_791() (interface{}, error) {
	return dm_build_1102.dm_build_892.dm_build_791()
}

const (
	Dm_build_1103 = Dm_build_697
)

type dm_build_1104 struct {
	dm_build_892
	dm_build_1105 []parameter
}

func dm_build_1106(dm_build_1107 *dm_build_410, dm_build_1108 *DmStatement, dm_build_1109 []parameter) *dm_build_1104 {
	dm_build_1110 := new(dm_build_1104)
	dm_build_1110.dm_build_809(dm_build_1107, Dm_build_693, dm_build_1108)
	dm_build_1110.dm_build_1105 = dm_build_1109
	return dm_build_1110
}

func (dm_build_1112 *dm_build_1104) dm_build_787() error {

	if dm_build_1112.dm_build_1105 == nil {
		dm_build_1112.dm_build_801.dm_build_413.Dm_build_265(Dm_build_1103, 0)
	} else {
		dm_build_1112.dm_build_801.dm_build_413.Dm_build_265(Dm_build_1103, int16(len(dm_build_1112.dm_build_1105)))
	}

	return dm_build_1112.dm_build_915(dm_build_1112.dm_build_1105)
}

type dm_build_1113 struct {
	dm_build_892
	dm_build_1114 bool
	dm_build_1115 int16
}

func dm_build_1116(dm_build_1117 *dm_build_410, dm_build_1118 *DmStatement, dm_build_1119 bool, dm_build_1120 int16) *dm_build_1113 {
	dm_build_1121 := new(dm_build_1113)
	dm_build_1121.dm_build_809(dm_build_1117, Dm_build_673, dm_build_1118)
	dm_build_1121.dm_build_1114 = dm_build_1119
	dm_build_1121.dm_build_1115 = dm_build_1120
	return dm_build_1121
}

func (dm_build_1123 *dm_build_1113) dm_build_787() error {

	dm_build_1124 := Dm_build_697

	if dm_build_1123.dm_build_801.dm_build_414.autoCommit {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 1)
	} else {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)
	}

	if dm_build_1123.dm_build_1114 {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 1)
	} else {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)
	}

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 1)

	if dm_build_1123.dm_build_801.dm_build_414.CompatibleOracle() {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)
	} else {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 1)
	}

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_265(dm_build_1124, dm_build_1123.dm_build_1115)

	if dm_build_1123.dm_build_804.maxRows <= 0 || dm_build_1123.dm_build_801.dm_build_414.dmConnector.enRsCache {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_273(dm_build_1124, INT64_MAX)
	} else {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_273(dm_build_1124, dm_build_1123.dm_build_804.maxRows)
	}

	if dm_build_1123.dm_build_801.dm_build_414.dmConnector.isBdtaRS {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, Dm_build_775)
	} else {
		dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, Dm_build_774)
	}

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_265(dm_build_1124, 0)

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 1)

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_261(dm_build_1124, 0)

	dm_build_1124 += dm_build_1123.dm_build_801.dm_build_413.Dm_build_269(dm_build_1124, dm_build_1123.dm_build_804.queryTimeout)

	dm_build_1123.dm_build_801.dm_build_413.Dm_build_191(dm_build_1123.dm_build_804.nativeSql, dm_build_1123.dm_build_801.dm_build_414.getServerEncoding(), dm_build_1123.dm_build_801.dm_build_414)

	return nil
}

func (dm_build_1126 *dm_build_1113) dm_build_791() (interface{}, error) {

	if dm_build_1126.dm_build_1114 {
		return dm_build_1126.dm_build_892.dm_build_791()
	}

	dm_build_1127 := NewExceInfo()
	dm_build_1128 := Dm_build_697

	dm_build_1127.retSqlType = dm_build_1126.dm_build_801.dm_build_413.Dm_build_344(dm_build_1128)
	dm_build_1128 += USINT_SIZE

	dm_build_1129 := dm_build_1126.dm_build_801.dm_build_413.Dm_build_344(dm_build_1128)
	dm_build_1128 += USINT_SIZE

	dm_build_1130 := dm_build_1126.dm_build_801.dm_build_413.Dm_build_344(dm_build_1128)
	dm_build_1128 += USINT_SIZE

	dm_build_1126.dm_build_801.dm_build_413.Dm_build_350(dm_build_1128)
	dm_build_1128 += DDWORD_SIZE

	dm_build_1126.dm_build_801.dm_build_414.TrxStatus = dm_build_1126.dm_build_801.dm_build_413.Dm_build_347(dm_build_1128)
	dm_build_1128 += ULINT_SIZE

	if dm_build_1129 > 0 {
		dm_build_1126.dm_build_804.params = dm_build_1126.dm_build_1131(int(dm_build_1129))
		dm_build_1126.dm_build_804.paramCount = dm_build_1129
	} else {
		dm_build_1126.dm_build_804.params = make([]parameter, 0)
		dm_build_1126.dm_build_804.paramCount = 0
	}

	if dm_build_1130 > 0 {
		dm_build_1126.dm_build_804.columns = dm_build_1126.dm_build_942(int(dm_build_1130), dm_build_1127.rsBdta)
	} else {
		dm_build_1126.dm_build_804.columns = make([]column, 0)
	}

	return dm_build_1127, nil
}

func (dm_build_1132 *dm_build_1113) dm_build_1131(dm_build_1133 int) []parameter {

	var dm_build_1134, dm_build_1135, dm_build_1136, dm_build_1137 int16

	dm_build_1138 := make([]parameter, dm_build_1133)
	for i := 0; i < dm_build_1133; i++ {

		dm_build_1138[i].InitParameter()

		dm_build_1138[i].colType = dm_build_1132.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_1138[i].prec = dm_build_1132.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_1138[i].scale = dm_build_1132.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_1138[i].nullable = dm_build_1132.dm_build_801.dm_build_413.Dm_build_203() != 0

		itemFlag := dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()

		if int(itemFlag)&Dm_build_891 != 0 {
			dm_build_1138[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_1138[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_1138[i].lob = int(itemFlag)&Dm_build_889 != 0

		dm_build_1132.dm_build_801.dm_build_413.Dm_build_203()

		dm_build_1138[i].ioType = byte(dm_build_1132.dm_build_801.dm_build_413.Dm_build_200())

		dm_build_1134 = dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_1135 = dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_1136 = dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()

		dm_build_1137 = dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()
		dm_build_1138[i].name = dm_build_1132.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_1134), dm_build_1132.dm_build_801.dm_build_414.getServerEncoding(), dm_build_1132.dm_build_801.dm_build_414)
		dm_build_1138[i].typeName = dm_build_1132.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_1135), dm_build_1132.dm_build_801.dm_build_414.getServerEncoding(), dm_build_1132.dm_build_801.dm_build_414)
		dm_build_1138[i].tableName = dm_build_1132.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_1136), dm_build_1132.dm_build_801.dm_build_414.getServerEncoding(), dm_build_1132.dm_build_801.dm_build_414)
		dm_build_1138[i].schemaName = dm_build_1132.dm_build_801.dm_build_413.Dm_build_240(int(dm_build_1137), dm_build_1132.dm_build_801.dm_build_414.getServerEncoding(), dm_build_1132.dm_build_801.dm_build_414)

		if dm_build_1138[i].lob {
			dm_build_1138[i].lobTabId = dm_build_1132.dm_build_801.dm_build_413.Dm_build_203()
			dm_build_1138[i].lobColId = dm_build_1132.dm_build_801.dm_build_413.Dm_build_200()
		}
	}

	for i := 0; i < dm_build_1133; i++ {

		if isComplexType(int(dm_build_1138[i].colType), int(dm_build_1138[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_1132.dm_build_801.dm_build_414)
			strDesc.unpack(dm_build_1132.dm_build_801.dm_build_413)
			dm_build_1138[i].typeDescriptor = strDesc
		}
	}

	return dm_build_1138
}

const (
	Dm_build_1139 = Dm_build_697
)

type dm_build_1140 struct {
	dm_build_800
	dm_build_1141 int16
	dm_build_1142 *Dm_build_0
	dm_build_1143 int32
}

func dm_build_1144(dm_build_1145 *dm_build_410, dm_build_1146 *DmStatement, dm_build_1147 int16, dm_build_1148 *Dm_build_0, dm_build_1149 int32) *dm_build_1140 {
	dm_build_1150 := new(dm_build_1140)
	dm_build_1150.dm_build_809(dm_build_1145, Dm_build_679, dm_build_1146)
	dm_build_1150.dm_build_1141 = dm_build_1147
	dm_build_1150.dm_build_1142 = dm_build_1148
	dm_build_1150.dm_build_1143 = dm_build_1149
	return dm_build_1150
}

func (dm_build_1152 *dm_build_1140) dm_build_787() error {
	dm_build_1152.dm_build_801.dm_build_413.Dm_build_265(Dm_build_1139, dm_build_1152.dm_build_1141)

	dm_build_1152.dm_build_801.dm_build_413.Dm_build_129(dm_build_1152.dm_build_1143)

	if dm_build_1152.dm_build_801.dm_build_414.NewLobFlag {
		dm_build_1152.dm_build_801.dm_build_413.Dm_build_129(-1)
	}
	dm_build_1152.dm_build_1142.Dm_build_7(dm_build_1152.dm_build_801.dm_build_413, int(dm_build_1152.dm_build_1143))
	return nil
}

type dm_build_1153 struct {
	dm_build_800
}

func dm_build_1154(dm_build_1155 *dm_build_410) *dm_build_1153 {
	dm_build_1156 := new(dm_build_1153)
	dm_build_1156.dm_build_805(dm_build_1155, Dm_build_677)
	return dm_build_1156
}

type dm_build_1157 struct {
	dm_build_800
	dm_build_1158 int32
}

func dm_build_1159(dm_build_1160 *dm_build_410, dm_build_1161 int32) *dm_build_1157 {
	dm_build_1162 := new(dm_build_1157)
	dm_build_1162.dm_build_805(dm_build_1160, Dm_build_690)
	dm_build_1162.dm_build_1158 = dm_build_1161
	return dm_build_1162
}

func (dm_build_1164 *dm_build_1157) dm_build_787() error {

	dm_build_1165 := Dm_build_697
	dm_build_1165 += dm_build_1164.dm_build_801.dm_build_413.Dm_build_269(dm_build_1165, g2dbIsoLevel(dm_build_1164.dm_build_1158))
	return nil
}

type dm_build_1166 struct {
	dm_build_800
	dm_build_1167 *lob
	dm_build_1168 byte
	dm_build_1169 int
	dm_build_1170 []byte
	dm_build_1171 int
	dm_build_1172 int
}

func dm_build_1173(dm_build_1174 *dm_build_410, dm_build_1175 *lob, dm_build_1176 byte, dm_build_1177 int, dm_build_1178 []byte,
	dm_build_1179 int, dm_build_1180 int) *dm_build_1166 {
	dm_build_1181 := new(dm_build_1166)
	dm_build_1181.dm_build_805(dm_build_1174, Dm_build_686)
	dm_build_1181.dm_build_1167 = dm_build_1175
	dm_build_1181.dm_build_1168 = dm_build_1176
	dm_build_1181.dm_build_1169 = dm_build_1177
	dm_build_1181.dm_build_1170 = dm_build_1178
	dm_build_1181.dm_build_1171 = dm_build_1179
	dm_build_1181.dm_build_1172 = dm_build_1180
	return dm_build_1181
}

func (dm_build_1183 *dm_build_1166) dm_build_787() error {

	dm_build_1183.dm_build_801.dm_build_413.Dm_build_121(byte(dm_build_1183.dm_build_1167.lobFlag))
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_121(dm_build_1183.dm_build_1168)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1183.dm_build_1167.blobId))
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.groupId)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.fileId)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(dm_build_1183.dm_build_1167.pageNo)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.curFileId)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(dm_build_1183.dm_build_1167.curPageNo)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(dm_build_1183.dm_build_1167.totalOffset)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(dm_build_1183.dm_build_1167.tabId)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.colId)
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_145(uint64(dm_build_1183.dm_build_1167.rowId))

	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(int32(dm_build_1183.dm_build_1169))
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(int32(dm_build_1183.dm_build_1172))
	dm_build_1183.dm_build_801.dm_build_413.Dm_build_157(dm_build_1183.dm_build_1170)

	if dm_build_1183.dm_build_801.dm_build_414.NewLobFlag {
		dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.exGroupId)
		dm_build_1183.dm_build_801.dm_build_413.Dm_build_125(dm_build_1183.dm_build_1167.exFileId)
		dm_build_1183.dm_build_801.dm_build_413.Dm_build_129(dm_build_1183.dm_build_1167.exPageNo)
	}
	return nil
}

func (dm_build_1185 *dm_build_1166) dm_build_791() (interface{}, error) {

	var dm_build_1186 = dm_build_1185.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1185.dm_build_1167.blobId = dm_build_1185.dm_build_801.dm_build_413.Dm_build_206()
	dm_build_1185.dm_build_1167.fileId = dm_build_1185.dm_build_801.dm_build_413.Dm_build_200()
	dm_build_1185.dm_build_1167.pageNo = dm_build_1185.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1185.dm_build_1167.curFileId = dm_build_1185.dm_build_801.dm_build_413.Dm_build_200()
	dm_build_1185.dm_build_1167.curPageNo = dm_build_1185.dm_build_801.dm_build_413.Dm_build_203()
	dm_build_1185.dm_build_1167.totalOffset = dm_build_1185.dm_build_801.dm_build_413.Dm_build_203()
	return dm_build_1186, nil
}

const (
	Dm_build_1187 = Dm_build_697

	Dm_build_1188 = Dm_build_1187 + ULINT_SIZE

	Dm_build_1189 = Dm_build_1188 + ULINT_SIZE

	Dm_build_1190 = Dm_build_1189 + BYTE_SIZE

	Dm_build_1191 = Dm_build_1190 + BYTE_SIZE

	Dm_build_1192 = Dm_build_1191 + BYTE_SIZE

	Dm_build_1193 = Dm_build_1192 + BYTE_SIZE

	Dm_build_1194 = Dm_build_697

	Dm_build_1195 = Dm_build_1194 + ULINT_SIZE

	Dm_build_1196 = Dm_build_1195 + ULINT_SIZE

	Dm_build_1197 = Dm_build_1196 + ULINT_SIZE

	Dm_build_1198 = Dm_build_1197 + ULINT_SIZE

	Dm_build_1199 = Dm_build_1198 + ULINT_SIZE

	Dm_build_1200 = Dm_build_1199 + BYTE_SIZE

	Dm_build_1201 = Dm_build_1200 + BYTE_SIZE

	Dm_build_1202 = Dm_build_1201 + BYTE_SIZE
)

type dm_build_1203 struct {
	dm_build_800
	dm_build_1204 *DmConnection
	dm_build_1205 int
	Dm_build_1206 int32
	Dm_build_1207 []byte
	dm_build_1208 byte
}

func dm_build_1209(dm_build_1210 *dm_build_410) *dm_build_1203 {
	dm_build_1211 := new(dm_build_1203)
	dm_build_1211.dm_build_805(dm_build_1210, Dm_build_695)
	dm_build_1211.dm_build_1204 = dm_build_1210.dm_build_414
	return dm_build_1211
}

func dm_build_1212(dm_build_1213 string, dm_build_1214 string) int {
	dm_build_1215 := strings.Split(dm_build_1213, ".")
	dm_build_1216 := strings.Split(dm_build_1214, ".")

	for i, serStr := range dm_build_1215 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_1216[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_1218 *dm_build_1203) dm_build_787() error {

	dm_build_1218.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1187, int32(0))
	dm_build_1218.dm_build_801.dm_build_413.Dm_build_269(Dm_build_1188, int32(dm_build_1218.dm_build_1204.dmConnector.compress))

	if dm_build_1218.dm_build_1204.dmConnector.loginEncrypt {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1190, 2)
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1189, 1)
	} else {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1190, 0)
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1189, 0)
	}

	if dm_build_1218.dm_build_1204.dmConnector.isBdtaRS {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1191, Dm_build_775)
	} else {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1191, Dm_build_774)
	}

	dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1192, byte(dm_build_1218.dm_build_1204.dmConnector.compressID))

	if dm_build_1218.dm_build_1204.dmConnector.loginCertificate != "" {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1193, 1)
	} else {
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_261(Dm_build_1193, 0)
	}

	dm_build_1219 := dm_build_1218.dm_build_1204.getServerEncoding()
	dm_build_1218.dm_build_801.dm_build_413.Dm_build_173(Dm_build_662, dm_build_1219, dm_build_1218.dm_build_801.dm_build_414)

	var dm_build_1220 byte
	if dm_build_1218.dm_build_1204.dmConnector.uKeyName != "" {
		dm_build_1220 = 1
	} else {
		dm_build_1220 = 0
	}

	dm_build_1218.dm_build_801.dm_build_413.Dm_build_121(0)

	if dm_build_1220 == 1 {

	}

	if dm_build_1218.dm_build_1204.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_1218.dm_build_801.dm_build_642()
		if err != nil {
			return err
		}
		dm_build_1218.dm_build_801.dm_build_413.Dm_build_161(clientPubKey)
	}

	return nil
}

func (dm_build_1222 *dm_build_1203) dm_build_790() error {
	dm_build_1222.dm_build_801.dm_build_413.Dm_build_95(0)
	dm_build_1222.dm_build_801.dm_build_413.Dm_build_103(Dm_build_696, false, true)
	return nil
}

func (dm_build_1224 *dm_build_1203) dm_build_791() (interface{}, error) {

	dm_build_1224.dm_build_1204.sslEncrypt = int(dm_build_1224.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1194))
	dm_build_1224.dm_build_1204.GlobalServerSeries = int(dm_build_1224.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1195))

	switch dm_build_1224.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1196) {
	case 1:
		dm_build_1224.dm_build_1204.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_1224.dm_build_1204.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_1224.dm_build_1204.serverEncoding = ENCODING_GB18030
	}

	dm_build_1224.dm_build_1204.dmConnector.compress = int(dm_build_1224.dm_build_801.dm_build_413.Dm_build_347(Dm_build_1197))
	dm_build_1225 := dm_build_1224.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1199)
	dm_build_1226 := dm_build_1224.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1200)
	dm_build_1224.dm_build_1204.dmConnector.isBdtaRS = dm_build_1224.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1201) > 0
	dm_build_1224.dm_build_1204.dmConnector.compressID = int8(dm_build_1224.dm_build_801.dm_build_413.Dm_build_341(Dm_build_1202))

	dm_build_1227 := dm_build_1224.dm_build_833()
	if dm_build_1227 != nil {
		return nil, dm_build_1227
	}

	dm_build_1228 := dm_build_1224.dm_build_801.dm_build_413.Dm_build_245(dm_build_1224.dm_build_1204.getServerEncoding(), dm_build_1224.dm_build_801.dm_build_414)
	if dm_build_1212(dm_build_1228, Dm_build_663) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_1224.dm_build_1204.ServerVersion = dm_build_1228
	dm_build_1224.dm_build_1204.Malini2 = dm_build_1212(dm_build_1228, Dm_build_664) > 0
	dm_build_1224.dm_build_1204.Execute2 = dm_build_1212(dm_build_1228, Dm_build_665) > 0
	dm_build_1224.dm_build_1204.LobEmptyCompOrcl = dm_build_1212(dm_build_1228, Dm_build_666) > 0

	if dm_build_1224.dm_build_801.dm_build_414.dmConnector.uKeyName != "" {
		dm_build_1224.dm_build_1208 = 1
	} else {
		dm_build_1224.dm_build_1208 = 0
	}

	if dm_build_1224.dm_build_1208 == 1 {
		dm_build_1224.dm_build_801.dm_build_419 = dm_build_1224.dm_build_801.dm_build_413.Dm_build_240(16, dm_build_1224.dm_build_1204.getServerEncoding(), dm_build_1224.dm_build_801.dm_build_414)
	}

	dm_build_1224.dm_build_1205 = -1
	dm_build_1229 := false
	dm_build_1230 := false
	dm_build_1224.Dm_build_1206 = -1
	if dm_build_1226 > 0 {
		dm_build_1224.dm_build_1205 = int(dm_build_1224.dm_build_801.dm_build_413.Dm_build_203())
	}

	if dm_build_1225 > 0 {

		if dm_build_1224.dm_build_1205 == -1 {
			dm_build_1229 = true
		} else {
			dm_build_1230 = true
		}

		dm_build_1224.Dm_build_1207 = dm_build_1224.dm_build_801.dm_build_413.Dm_build_228()
	}

	if dm_build_1226 == 2 {
		dm_build_1224.Dm_build_1206 = dm_build_1224.dm_build_801.dm_build_413.Dm_build_203()
	}
	dm_build_1224.dm_build_801.dm_build_416 = dm_build_1229
	dm_build_1224.dm_build_801.dm_build_417 = dm_build_1230

	return nil, nil
}

type dm_build_1231 struct {
	dm_build_800
}

func dm_build_1232(dm_build_1233 *dm_build_410, dm_build_1234 *DmStatement) *dm_build_1231 {
	dm_build_1235 := new(dm_build_1231)
	dm_build_1235.dm_build_809(dm_build_1233, Dm_build_671, dm_build_1234)
	return dm_build_1235
}

func (dm_build_1237 *dm_build_1231) dm_build_787() error {

	dm_build_1237.dm_build_801.dm_build_413.Dm_build_261(Dm_build_697, 1)
	return nil
}

func (dm_build_1239 *dm_build_1231) dm_build_791() (interface{}, error) {

	dm_build_1239.dm_build_804.id = dm_build_1239.dm_build_801.dm_build_413.Dm_build_347(Dm_build_698)

	dm_build_1239.dm_build_804.readBaseColName = dm_build_1239.dm_build_801.dm_build_413.Dm_build_341(Dm_build_697) == 1
	return nil, nil
}

type dm_build_1240 struct {
	dm_build_800
	dm_build_1241 int32
}

func dm_build_1242(dm_build_1243 *dm_build_410, dm_build_1244 int32) *dm_build_1240 {
	dm_build_1245 := new(dm_build_1240)
	dm_build_1245.dm_build_805(dm_build_1243, Dm_build_672)
	dm_build_1245.dm_build_1241 = dm_build_1244
	return dm_build_1245
}

func (dm_build_1247 *dm_build_1240) dm_build_788() {
	dm_build_1247.dm_build_800.dm_build_788()
	dm_build_1247.dm_build_801.dm_build_413.Dm_build_269(Dm_build_698, dm_build_1247.dm_build_1241)
}

type dm_build_1248 struct {
	dm_build_800
	dm_build_1249 []uint32
}

func dm_build_1250(dm_build_1251 *dm_build_410, dm_build_1252 []uint32) *dm_build_1248 {
	dm_build_1253 := new(dm_build_1248)
	dm_build_1253.dm_build_805(dm_build_1251, Dm_build_692)
	dm_build_1253.dm_build_1249 = dm_build_1252
	return dm_build_1253
}

func (dm_build_1255 *dm_build_1248) dm_build_787() error {

	dm_build_1255.dm_build_801.dm_build_413.Dm_build_289(Dm_build_697, uint16(len(dm_build_1255.dm_build_1249)))

	for _, tableID := range dm_build_1255.dm_build_1249 {
		dm_build_1255.dm_build_801.dm_build_413.Dm_build_141(uint32(tableID))
	}

	return nil
}

func (dm_build_1257 *dm_build_1248) dm_build_791() (interface{}, error) {
	dm_build_1258 := dm_build_1257.dm_build_801.dm_build_413.Dm_build_362(Dm_build_697)
	if dm_build_1258 <= 0 {
		return nil, nil
	}

	dm_build_1259 := make([]int64, dm_build_1258)
	for i := 0; i < int(dm_build_1258); i++ {
		dm_build_1259[i] = dm_build_1257.dm_build_801.dm_build_413.Dm_build_206()
	}

	return dm_build_1259, nil
}
