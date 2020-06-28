/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"crypto/tls"
	"dm/security"
	"net"
	"strconv"
	"time"
)

const (
	Dm_build_408 = 8192
	Dm_build_409 = 2 * time.Second
)

type dm_build_410 struct {
	dm_build_411 *net.TCPConn
	dm_build_412 *tls.Conn
	dm_build_413 *Dm_build_78
	dm_build_414 *DmConnection
	dm_build_415 security.Cipher
	dm_build_416 bool
	dm_build_417 bool
	dm_build_418 *security.DhKey
	dm_build_419 string
	dm_build_420 bool
}

func dm_build_421(dm_build_422 *DmConnection) (*dm_build_410, error) {
	dm_build_423, dm_build_424 := dm_build_426(dm_build_422.dmConnector.host+":"+strconv.Itoa(dm_build_422.dmConnector.port), time.Duration(dm_build_422.dmConnector.socketTimeout)*time.Second)
	if dm_build_424 != nil {
		return nil, dm_build_424
	}

	dm_build_425 := dm_build_410{}
	dm_build_425.dm_build_411 = dm_build_423
	dm_build_425.dm_build_413 = Dm_build_81(Dm_build_668)
	dm_build_425.dm_build_414 = dm_build_422
	dm_build_425.dm_build_416 = false
	dm_build_425.dm_build_417 = false
	dm_build_425.dm_build_419 = ""
	dm_build_425.dm_build_420 = false
	dm_build_422.Access = &dm_build_425

	return &dm_build_425, nil
}

func dm_build_426(dm_build_427 string, dm_build_428 time.Duration) (*net.TCPConn, error) {
	dm_build_429, dm_build_430 := net.DialTimeout("tcp", dm_build_427, dm_build_428)
	if dm_build_430 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_427).throw()
	}

	if tcpConn, ok := dm_build_429.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_409)
		tcpConn.SetNoDelay(true)

		tcpConn.SetReadBuffer(Dm_build_408)
		tcpConn.SetWriteBuffer(Dm_build_408)
		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_432 *dm_build_410) dm_build_431(dm_build_433 dm_build_785) bool {
	var dm_build_434 = dm_build_432.dm_build_414.dmConnector.compress
	if dm_build_433.dm_build_799() == Dm_build_695 || dm_build_434 == Dm_build_743 {
		return false
	}

	if dm_build_434 == Dm_build_741 {
		return true
	} else if dm_build_434 == Dm_build_742 {
		return !dm_build_432.dm_build_414.Local && dm_build_433.dm_build_797() > Dm_build_740
	}

	return false
}

func (dm_build_436 *dm_build_410) dm_build_435(dm_build_437 dm_build_785) bool {
	var dm_build_438 = dm_build_436.dm_build_414.dmConnector.compress
	if dm_build_437.dm_build_799() == Dm_build_695 || dm_build_438 == Dm_build_743 {
		return false
	}

	if dm_build_438 == Dm_build_741 {
		return true
	} else if dm_build_438 == Dm_build_742 {
		return dm_build_436.dm_build_413.Dm_build_341(Dm_build_703) == 1
	}

	return false
}

func (dm_build_440 *dm_build_410) dm_build_439(dm_build_441 dm_build_785) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_443 := dm_build_441.dm_build_797()

	if dm_build_443 > 0 {

		if dm_build_440.dm_build_431(dm_build_441) {
			var retBytes, err = Compress(dm_build_440.dm_build_413, Dm_build_696, int(dm_build_443), int(dm_build_440.dm_build_414.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_440.dm_build_413.Dm_build_92(Dm_build_696)

			dm_build_440.dm_build_413.Dm_build_129(dm_build_443)

			dm_build_440.dm_build_413.Dm_build_157(retBytes)

			dm_build_441.dm_build_798(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_440.dm_build_413.Dm_build_261(Dm_build_703, 1)
		}

		if dm_build_440.dm_build_417 {
			dm_build_443 = dm_build_441.dm_build_797()
			var retBytes = dm_build_440.dm_build_415.Encrypt(dm_build_440.dm_build_413.Dm_build_368(Dm_build_696, int(dm_build_443)), true)

			dm_build_440.dm_build_413.Dm_build_92(Dm_build_696)

			dm_build_440.dm_build_413.Dm_build_157(retBytes)

			dm_build_441.dm_build_798(int32(len(retBytes)))
		}
	}

	dm_build_441.dm_build_794()
	if dm_build_440.dm_build_659(dm_build_441) {
		if dm_build_440.dm_build_412 != nil {
			dm_build_440.dm_build_413.Dm_build_95(0)
			dm_build_440.dm_build_413.Dm_build_114(dm_build_440.dm_build_412)
		}
	} else {
		dm_build_440.dm_build_413.Dm_build_95(0)
		dm_build_440.dm_build_413.Dm_build_114(dm_build_440.dm_build_411)
	}
	return nil
}

func (dm_build_445 *dm_build_410) dm_build_444(dm_build_446 dm_build_785) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_448 := int32(0)
	if dm_build_445.dm_build_659(dm_build_446) {
		if dm_build_445.dm_build_412 != nil {
			dm_build_445.dm_build_413.Dm_build_92(0)
			dm_build_445.dm_build_413.Dm_build_108(dm_build_445.dm_build_412, Dm_build_696)
			dm_build_448 = dm_build_446.dm_build_797()
			if dm_build_448 > 0 {
				dm_build_445.dm_build_413.Dm_build_108(dm_build_445.dm_build_412, int(dm_build_448))
			}
		}
	} else {

		dm_build_445.dm_build_413.Dm_build_92(0)
		dm_build_445.dm_build_413.Dm_build_108(dm_build_445.dm_build_411, Dm_build_696)
		dm_build_448 = dm_build_446.dm_build_797()

		if dm_build_448 > 0 {
			dm_build_445.dm_build_413.Dm_build_108(dm_build_445.dm_build_411, int(dm_build_448))
		}
	}

	dm_build_446.dm_build_795()

	if dm_build_448 <= 0 {
		return nil
	}

	if dm_build_445.dm_build_417 {
		dm_build_448 = dm_build_446.dm_build_797()
		ebytes := dm_build_445.dm_build_413.Dm_build_368(Dm_build_696, int(dm_build_448))
		bytes, err := dm_build_445.dm_build_415.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_445.dm_build_413.Dm_build_92(Dm_build_696)
		dm_build_445.dm_build_413.Dm_build_157(bytes)
		dm_build_446.dm_build_798(int32(len(bytes)))
	}

	if dm_build_445.dm_build_435(dm_build_446) {

		dm_build_448 = dm_build_446.dm_build_797()
		cbytes := dm_build_445.dm_build_413.Dm_build_368(Dm_build_696+ULINT_SIZE, int(dm_build_448-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_445.dm_build_414.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_445.dm_build_413.Dm_build_92(Dm_build_696)
		dm_build_445.dm_build_413.Dm_build_157(bytes)
		dm_build_446.dm_build_798(int32(len(bytes)))
	}
	return nil
}

func (dm_build_450 *dm_build_410) dm_build_449(dm_build_451 dm_build_785) (dm_build_452 interface{}, dm_build_453 error) {
	dm_build_453 = dm_build_451.dm_build_789(dm_build_451)
	if dm_build_453 != nil {
		return nil, dm_build_453
	}

	dm_build_453 = dm_build_450.dm_build_439(dm_build_451)
	if dm_build_453 != nil {
		return nil, dm_build_453
	}

	dm_build_453 = dm_build_450.dm_build_444(dm_build_451)
	if dm_build_453 != nil {
		return nil, dm_build_453
	}

	return dm_build_451.dm_build_793(dm_build_451)
}

func (dm_build_455 *dm_build_410) dm_build_454() (*dm_build_1203, error) {

	Dm_build_456 := dm_build_1209(dm_build_455)
	_, dm_build_457 := dm_build_455.dm_build_449(Dm_build_456)
	if dm_build_457 != nil {
		return nil, dm_build_457
	}

	return Dm_build_456, nil
}

func (dm_build_459 *dm_build_410) dm_build_458() error {

	dm_build_460 := dm_build_1080(dm_build_459)
	_, dm_build_461 := dm_build_459.dm_build_449(dm_build_460)
	if dm_build_461 != nil {
		return dm_build_461
	}

	return nil
}

func (dm_build_463 *dm_build_410) dm_build_462() error {

	var dm_build_464 *dm_build_1203
	var err error
	if dm_build_464, err = dm_build_463.dm_build_454(); err != nil {
		return err
	}

	if dm_build_463.dm_build_414.sslEncrypt == 2 {
		if err = dm_build_463.dm_build_655(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_463.dm_build_414.sslEncrypt == 1 {
		if err = dm_build_463.dm_build_655(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_463.dm_build_417 || dm_build_463.dm_build_416 {
		k, err := dm_build_463.dm_build_645()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_464.Dm_build_1207)
		encryptType := dm_build_464.dm_build_1205
		hashType := int(dm_build_464.Dm_build_1206)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_463.dm_build_648(encryptType, sessionKey, dm_build_463.dm_build_414.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_463.dm_build_458(); err != nil {
		return err
	}
	return nil
}

func (dm_build_467 *dm_build_410) Dm_build_466(dm_build_468 *DmStatement) error {
	dm_build_469 := dm_build_1232(dm_build_467, dm_build_468)
	_, dm_build_470 := dm_build_467.dm_build_449(dm_build_469)
	if dm_build_470 != nil {
		return dm_build_470
	}

	return nil
}

func (dm_build_472 *dm_build_410) Dm_build_471(dm_build_473 int32) error {
	dm_build_474 := dm_build_1242(dm_build_472, dm_build_473)
	_, dm_build_475 := dm_build_472.dm_build_449(dm_build_474)
	if dm_build_475 != nil {
		return dm_build_475
	}

	return nil
}

func (dm_build_477 *dm_build_410) Dm_build_476(dm_build_478 *DmStatement, dm_build_479 bool, dm_build_480 int16) (*execInfo, error) {
	dm_build_481 := dm_build_1116(dm_build_477, dm_build_478, dm_build_479, dm_build_480)
	dm_build_482, dm_build_483 := dm_build_477.dm_build_449(dm_build_481)
	if dm_build_483 != nil {
		return nil, dm_build_483
	}
	return dm_build_482.(*execInfo), nil
}

func (dm_build_485 *dm_build_410) Dm_build_484(dm_build_486 *DmStatement, dm_build_487 int16) (*execInfo, error) {
	return dm_build_485.Dm_build_476(dm_build_486, false, Dm_build_747)
}

func (dm_build_489 *dm_build_410) Dm_build_488(dm_build_490 *DmStatement, dm_build_491 []OptParameter) (*execInfo, error) {
	dm_build_492, dm_build_493 := dm_build_489.dm_build_449(dm_build_878(dm_build_489, dm_build_490, dm_build_491))
	if dm_build_493 != nil {
		return nil, dm_build_493
	}

	return dm_build_492.(*execInfo), nil
}

func (dm_build_495 *dm_build_410) Dm_build_494(dm_build_496 *DmStatement, dm_build_497 int16) (*execInfo, error) {
	return dm_build_495.Dm_build_476(dm_build_496, true, dm_build_497)
}

func (dm_build_499 *dm_build_410) Dm_build_498(dm_build_500 *DmStatement, dm_build_501 [][]interface{}) (*execInfo, error) {
	dm_build_502 := dm_build_901(dm_build_499, dm_build_500, dm_build_501)
	dm_build_503, dm_build_504 := dm_build_499.dm_build_449(dm_build_502)
	if dm_build_504 != nil {
		return nil, dm_build_504
	}
	return dm_build_503.(*execInfo), nil
}

func (dm_build_506 *dm_build_410) Dm_build_505(dm_build_507 *DmStatement, dm_build_508 [][]interface{}) (*execInfo, error) {
	var dm_build_509, dm_build_510 = 0, 0
	var dm_build_511 = len(dm_build_508)
	var dm_build_512 [][]interface{}
	var dm_build_513 = NewExceInfo()
	dm_build_513.updateCounts = make([]int64, dm_build_511)
	var dm_build_514 = false
	var dm_build_515 = false
	for dm_build_509 < dm_build_511 {
		for dm_build_510 = dm_build_509; dm_build_510 < dm_build_511; dm_build_510++ {
			paramData := dm_build_508[dm_build_510]
			bindData := make([]interface{}, dm_build_507.paramCount)
			dm_build_514 = false
			for icol := 0; icol < int(dm_build_507.paramCount); icol++ {
				if dm_build_507.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_506.dm_build_628(bindData, paramData, icol) {
					dm_build_514 = true
					break
				}
			}

			if dm_build_514 {
				break
			}
			dm_build_512 = append(dm_build_512, bindData)
		}

		if dm_build_510 != dm_build_509 {
			tmpExecInfo, err := dm_build_506.Dm_build_498(dm_build_507, dm_build_512)
			if err != nil {
				return nil, err
			}
			dm_build_512 = dm_build_512[0:0]
			if dm_build_510-dm_build_509 == 1 {
				dm_build_513.updateCounts[dm_build_509] = tmpExecInfo.updateCount
			} else if tmpExecInfo.updateCounts != nil {
				copy(dm_build_513.updateCounts[dm_build_509:dm_build_510], tmpExecInfo.updateCounts[0:dm_build_510-dm_build_509])
			}
		}

		if dm_build_510 < dm_build_511 {
			tmpExecInfo, err := dm_build_506.Dm_build_516(dm_build_507, dm_build_508[dm_build_510], dm_build_515)
			if err != nil {
				return nil, err
			}

			dm_build_515 = true
			dm_build_513.updateCounts[dm_build_510] = tmpExecInfo.updateCount
		}

		dm_build_509 = dm_build_510 + 1
	}
	for _, i := range dm_build_513.updateCounts {
		dm_build_513.updateCount += i
	}
	return dm_build_513, nil
}

func (dm_build_517 *dm_build_410) Dm_build_516(dm_build_518 *DmStatement, dm_build_519 []interface{}, dm_build_520 bool) (*execInfo, error) {

	var dm_build_521 = make([]interface{}, dm_build_518.paramCount)
	for icol := 0; icol < int(dm_build_518.paramCount); icol++ {
		if dm_build_518.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_517.dm_build_628(dm_build_521, dm_build_519, icol) {

			if !dm_build_520 {
				preExecute := dm_build_1106(dm_build_517, dm_build_518, dm_build_518.params)
				dm_build_517.dm_build_449(preExecute)
				dm_build_520 = true
			}

			dm_build_517.dm_build_634(dm_build_518, dm_build_518.params[icol], icol, dm_build_519[icol].(iOffRowBinder))
			dm_build_521[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_522 = make([][]interface{}, 1, 1)
	dm_build_522[0] = dm_build_521

	dm_build_523 := dm_build_901(dm_build_517, dm_build_518, dm_build_522)
	dm_build_524, dm_build_525 := dm_build_517.dm_build_449(dm_build_523)
	if dm_build_525 != nil {
		return nil, dm_build_525
	}
	return dm_build_524.(*execInfo), nil
}

func (dm_build_527 *dm_build_410) Dm_build_526(dm_build_528 *DmStatement, dm_build_529 int16) (*execInfo, error) {
	dm_build_530 := dm_build_1094(dm_build_527, dm_build_528, dm_build_529)

	dm_build_531, dm_build_532 := dm_build_527.dm_build_449(dm_build_530)
	if dm_build_532 != nil {
		return nil, dm_build_532
	}
	return dm_build_531.(*execInfo), nil
}

func (dm_build_534 *dm_build_410) Dm_build_533(dm_build_535 *innerRows, dm_build_536 int64) (*execInfo, error) {
	dm_build_537 := dm_build_1001(dm_build_534, dm_build_535, dm_build_536, INT64_MAX)
	dm_build_538, dm_build_539 := dm_build_534.dm_build_449(dm_build_537)
	if dm_build_539 != nil {
		return nil, dm_build_539
	}
	return dm_build_538.(*execInfo), nil
}

func (dm_build_541 *dm_build_410) Commit() error {
	dm_build_542 := dm_build_864(dm_build_541)
	_, dm_build_543 := dm_build_541.dm_build_449(dm_build_542)
	if dm_build_543 != nil {
		return dm_build_543
	}

	return nil
}

func (dm_build_545 *dm_build_410) Rollback() error {
	dm_build_546 := dm_build_1154(dm_build_545)
	_, dm_build_547 := dm_build_545.dm_build_449(dm_build_546)
	if dm_build_547 != nil {
		return dm_build_547
	}

	return nil
}

func (dm_build_549 *dm_build_410) Dm_build_548(dm_build_550 *DmConnection) error {
	dm_build_551 := dm_build_1159(dm_build_549, dm_build_550.IsoLevel)
	_, dm_build_552 := dm_build_549.dm_build_449(dm_build_551)
	if dm_build_552 != nil {
		return dm_build_552
	}

	return nil
}

func (dm_build_554 *dm_build_410) Dm_build_553(dm_build_555 *DmStatement, dm_build_556 string) error {
	dm_build_557 := dm_build_869(dm_build_554, dm_build_555, dm_build_556)
	_, dm_build_558 := dm_build_554.dm_build_449(dm_build_557)
	if dm_build_558 != nil {
		return dm_build_558
	}

	return nil
}

func (dm_build_560 *dm_build_410) Dm_build_559(dm_build_561 []uint32) ([]int64, error) {
	dm_build_562 := dm_build_1250(dm_build_560, dm_build_561)
	dm_build_563, dm_build_564 := dm_build_560.dm_build_449(dm_build_562)
	if dm_build_564 != nil {
		return nil, dm_build_564
	}
	return dm_build_563.([]int64), nil
}

func (dm_build_566 *dm_build_410) Close() error {
	if dm_build_566.dm_build_420 {
		return nil
	}

	dm_build_567 := dm_build_566.dm_build_411.Close()
	if dm_build_567 != nil {
		return dm_build_567
	}

	dm_build_566.dm_build_414 = nil
	dm_build_566.dm_build_420 = true
	return nil
}

func (dm_build_569 *dm_build_410) dm_build_568(dm_build_570 *lob) (int64, error) {
	dm_build_571 := dm_build_1032(dm_build_569, dm_build_570)
	dm_build_572, dm_build_573 := dm_build_569.dm_build_449(dm_build_571)
	if dm_build_573 != nil {
		return 0, dm_build_573
	}
	return dm_build_572.(int64), nil
}

func (dm_build_575 *dm_build_410) dm_build_574(dm_build_576 *DmBlob, dm_build_577 int32, dm_build_578 int32) ([]byte, error) {

	dm_build_579 := dm_build_1019(dm_build_575, &dm_build_576.lob, int(dm_build_577), int(dm_build_578))
	dm_build_580, dm_build_581 := dm_build_575.dm_build_449(dm_build_579)
	if dm_build_581 != nil {
		return nil, dm_build_581
	}
	return dm_build_580.([]byte), nil
}

func (dm_build_583 *dm_build_410) dm_build_582(dm_build_584 *DmClob, dm_build_585 int32, dm_build_586 int32) (string, error) {

	dm_build_587 := dm_build_1019(dm_build_583, &dm_build_584.lob, int(dm_build_585), int(dm_build_586))
	dm_build_588, dm_build_589 := dm_build_583.dm_build_449(dm_build_587)
	if dm_build_589 != nil {
		return "", dm_build_589
	}
	dm_build_590 := dm_build_588.([]byte)
	return Dm_build_1261.Dm_build_1415(dm_build_590, 0, len(dm_build_590), dm_build_584.serverEncoding, dm_build_583.dm_build_414), nil
}

func (dm_build_592 *dm_build_410) dm_build_591(dm_build_593 *DmClob, dm_build_594 int, dm_build_595 string, dm_build_596 string) (int, error) {
	var dm_build_597 = Dm_build_1261.Dm_build_1471(dm_build_595, dm_build_596, dm_build_592.dm_build_414)
	var dm_build_598 = 0
	var dm_build_599 = len(dm_build_597)
	var dm_build_600 = 0
	var dm_build_601 = 0
	var dm_build_602 = 0
	var dm_build_603 = dm_build_599/Dm_build_779 + 1
	var dm_build_604 byte = 0
	var dm_build_605 byte = 0x01
	var dm_build_606 byte = 0x02
	for i := 0; i < dm_build_603; i++ {
		dm_build_604 = 0
		if i == 0 {
			dm_build_604 |= dm_build_605
		}
		if i == dm_build_603-1 {
			dm_build_604 |= dm_build_606
		}
		dm_build_602 = dm_build_599 - dm_build_601
		if dm_build_602 > Dm_build_779 {
			dm_build_602 = Dm_build_779
		}

		setLobData := dm_build_1173(dm_build_592, &dm_build_593.lob, dm_build_604, dm_build_594, dm_build_597, dm_build_598, dm_build_602)
		ret, err := dm_build_592.dm_build_449(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_600, nil
		} else {
			dm_build_594 += int(tmp)
			dm_build_600 += int(tmp)
			dm_build_601 += dm_build_602
			dm_build_598 += dm_build_602
		}
	}
	return dm_build_600, nil
}

func (dm_build_608 *dm_build_410) dm_build_607(dm_build_609 *DmBlob, dm_build_610 int, dm_build_611 []byte) (int, error) {
	var dm_build_612 = 0
	var dm_build_613 = len(dm_build_611)
	var dm_build_614 = 0
	var dm_build_615 = 0
	var dm_build_616 = 0
	var dm_build_617 = dm_build_613/Dm_build_779 + 1
	var dm_build_618 byte = 0
	var dm_build_619 byte = 0x01
	var dm_build_620 byte = 0x02
	for i := 0; i < dm_build_617; i++ {
		dm_build_618 = 0
		if i == 0 {
			dm_build_618 |= dm_build_619
		}
		if i == dm_build_617-1 {
			dm_build_618 |= dm_build_620
		}
		dm_build_616 = dm_build_613 - dm_build_615
		if dm_build_616 > Dm_build_779 {
			dm_build_616 = Dm_build_779
		}

		setLobData := dm_build_1173(dm_build_608, &dm_build_609.lob, dm_build_618, dm_build_610, dm_build_611, dm_build_612, dm_build_616)
		ret, err := dm_build_608.dm_build_449(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_614, nil
		} else {
			dm_build_610 += int(tmp)
			dm_build_614 += int(tmp)
			dm_build_615 += dm_build_616
			dm_build_612 += dm_build_616
		}
	}
	return dm_build_614, nil
}

func (dm_build_622 *dm_build_410) dm_build_621(dm_build_623 *lob, dm_build_624 int) (int64, error) {
	dm_build_625 := dm_build_1043(dm_build_622, dm_build_623, dm_build_624)
	dm_build_626, dm_build_627 := dm_build_622.dm_build_449(dm_build_625)
	if dm_build_627 != nil {
		return dm_build_623.length, dm_build_627
	}
	return dm_build_626.(int64), nil
}

func (dm_build_629 *dm_build_410) dm_build_628(dm_build_630 []interface{}, dm_build_631 []interface{}, dm_build_632 int) bool {
	var dm_build_633 = false
	if dm_build_632 >= len(dm_build_631) || dm_build_631[dm_build_632] == nil {
		dm_build_630[dm_build_632] = ParamDataEnum_Null
	} else if binder, ok := dm_build_631[dm_build_632].(iOffRowBinder); ok {
		dm_build_633 = true
		dm_build_630[dm_build_632] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_629.dm_build_414) {
			dm_build_630[dm_build_632] = &lobCtl{lob.buildCtlData()}
			dm_build_633 = false
		}
	} else {
		dm_build_630[dm_build_632] = dm_build_631[dm_build_632]
	}
	return dm_build_633
}

func (dm_build_635 *dm_build_410) dm_build_634(dm_build_636 *DmStatement, dm_build_637 parameter, dm_build_638 int, dm_build_639 iOffRowBinder) error {
	var dm_build_640 = Dm_build_4()
	dm_build_639.read(dm_build_640)
	var dm_build_641 = 0
	for !dm_build_639.isReadOver() || dm_build_640.Dm_build_5() > 0 {
		if !dm_build_639.isReadOver() && dm_build_640.Dm_build_5() < Dm_build_779 {
			dm_build_639.read(dm_build_640)
		}
		if dm_build_640.Dm_build_5() > Dm_build_779 {
			dm_build_641 = Dm_build_779
		} else {
			dm_build_641 = dm_build_640.Dm_build_5()
		}

		putData := dm_build_1144(dm_build_635, dm_build_636, int16(dm_build_638), dm_build_640, int32(dm_build_641))
		_, err := dm_build_635.dm_build_449(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_643 *dm_build_410) dm_build_642() ([]byte, error) {
	var dm_build_644 error
	if dm_build_643.dm_build_418 == nil {
		if dm_build_643.dm_build_418, dm_build_644 = security.NewClientKeyPair(); dm_build_644 != nil {
			return nil, dm_build_644
		}
	}
	return security.Bn2Bytes(dm_build_643.dm_build_418.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_646 *dm_build_410) dm_build_645() (*security.DhKey, error) {
	var dm_build_647 error
	if dm_build_646.dm_build_418 == nil {
		if dm_build_646.dm_build_418, dm_build_647 = security.NewClientKeyPair(); dm_build_647 != nil {
			return nil, dm_build_647
		}
	}
	return dm_build_646.dm_build_418, nil
}

func (dm_build_649 *dm_build_410) dm_build_648(dm_build_650 int, dm_build_651 []byte, dm_build_652 string, dm_build_653 int) (dm_build_654 error) {
	if dm_build_650 > 0 && dm_build_650 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_651 != nil {
		dm_build_649.dm_build_415, dm_build_654 = security.NewSymmCipher(dm_build_650, dm_build_651)
	} else if dm_build_650 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_649.dm_build_415, dm_build_654 = security.NewThirdPartCipher(dm_build_650, dm_build_651, dm_build_652, dm_build_653); dm_build_654 != nil {
			dm_build_654 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_654.Error()).throw()
		}
	}
	return
}

func (dm_build_656 *dm_build_410) dm_build_655(dm_build_657 bool) (dm_build_658 error) {
	if dm_build_656.dm_build_412, dm_build_658 = security.NewTLSFromTCP(dm_build_656.dm_build_411, dm_build_656.dm_build_414.dmConnector.sslCertPath, dm_build_656.dm_build_414.dmConnector.sslKeyPath, dm_build_656.dm_build_414.dmConnector.user); dm_build_658 != nil {
		return
	}
	if !dm_build_657 {
		dm_build_656.dm_build_412 = nil
	}
	return
}

func (dm_build_660 *dm_build_410) dm_build_659(dm_build_661 dm_build_785) bool {
	return dm_build_661.dm_build_799() != Dm_build_695 && dm_build_660.dm_build_414.sslEncrypt == 1
}
