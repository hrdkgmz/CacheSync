/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1260 struct{}

var Dm_build_1261 = &dm_build_1260{}

func (Dm_build_1263 *dm_build_1260) Dm_build_1262(dm_build_1264 []byte, dm_build_1265 int, dm_build_1266 byte) int {
	dm_build_1264[dm_build_1265] = dm_build_1266
	return 1
}

func (Dm_build_1268 *dm_build_1260) Dm_build_1267(dm_build_1269 []byte, dm_build_1270 int, dm_build_1271 int8) int {
	dm_build_1269[dm_build_1270] = byte(dm_build_1271)
	return 1
}

func (Dm_build_1273 *dm_build_1260) Dm_build_1272(dm_build_1274 []byte, dm_build_1275 int, dm_build_1276 int16) int {
	dm_build_1274[dm_build_1275] = byte(dm_build_1276)
	dm_build_1275++
	dm_build_1274[dm_build_1275] = byte(dm_build_1276 >> 8)
	return 2
}

func (Dm_build_1278 *dm_build_1260) Dm_build_1277(dm_build_1279 []byte, dm_build_1280 int, dm_build_1281 int32) int {
	dm_build_1279[dm_build_1280] = byte(dm_build_1281)
	dm_build_1280++
	dm_build_1279[dm_build_1280] = byte(dm_build_1281 >> 8)
	dm_build_1280++
	dm_build_1279[dm_build_1280] = byte(dm_build_1281 >> 16)
	dm_build_1280++
	dm_build_1279[dm_build_1280] = byte(dm_build_1281 >> 24)
	dm_build_1280++
	return 4
}

func (Dm_build_1283 *dm_build_1260) Dm_build_1282(dm_build_1284 []byte, dm_build_1285 int, dm_build_1286 int64) int {
	dm_build_1284[dm_build_1285] = byte(dm_build_1286)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 8)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 16)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 24)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 32)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 40)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 48)
	dm_build_1285++
	dm_build_1284[dm_build_1285] = byte(dm_build_1286 >> 56)
	return 8
}

func (Dm_build_1288 *dm_build_1260) Dm_build_1287(dm_build_1289 []byte, dm_build_1290 int, dm_build_1291 float32) int {
	return Dm_build_1288.Dm_build_1307(dm_build_1289, dm_build_1290, math.Float32bits(dm_build_1291))
}

func (Dm_build_1293 *dm_build_1260) Dm_build_1292(dm_build_1294 []byte, dm_build_1295 int, dm_build_1296 float64) int {
	return Dm_build_1293.Dm_build_1312(dm_build_1294, dm_build_1295, math.Float64bits(dm_build_1296))
}

func (Dm_build_1298 *dm_build_1260) Dm_build_1297(dm_build_1299 []byte, dm_build_1300 int, dm_build_1301 uint8) int {
	dm_build_1299[dm_build_1300] = byte(dm_build_1301)
	return 1
}

func (Dm_build_1303 *dm_build_1260) Dm_build_1302(dm_build_1304 []byte, dm_build_1305 int, dm_build_1306 uint16) int {
	dm_build_1304[dm_build_1305] = byte(dm_build_1306)
	dm_build_1305++
	dm_build_1304[dm_build_1305] = byte(dm_build_1306 >> 8)
	return 2
}

func (Dm_build_1308 *dm_build_1260) Dm_build_1307(dm_build_1309 []byte, dm_build_1310 int, dm_build_1311 uint32) int {
	dm_build_1309[dm_build_1310] = byte(dm_build_1311)
	dm_build_1310++
	dm_build_1309[dm_build_1310] = byte(dm_build_1311 >> 8)
	dm_build_1310++
	dm_build_1309[dm_build_1310] = byte(dm_build_1311 >> 16)
	dm_build_1310++
	dm_build_1309[dm_build_1310] = byte(dm_build_1311 >> 24)
	return 3
}

func (Dm_build_1313 *dm_build_1260) Dm_build_1312(dm_build_1314 []byte, dm_build_1315 int, dm_build_1316 uint64) int {
	dm_build_1314[dm_build_1315] = byte(dm_build_1316)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 8)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 16)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 24)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 32)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 40)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 48)
	dm_build_1315++
	dm_build_1314[dm_build_1315] = byte(dm_build_1316 >> 56)
	return 3
}

func (Dm_build_1318 *dm_build_1260) Dm_build_1317(dm_build_1319 []byte, dm_build_1320 int, dm_build_1321 []byte, dm_build_1322 int, dm_build_1323 int) int {
	copy(dm_build_1319[dm_build_1320:dm_build_1320+dm_build_1323], dm_build_1321[dm_build_1322:dm_build_1322+dm_build_1323])
	return dm_build_1323
}

func (Dm_build_1325 *dm_build_1260) Dm_build_1324(dm_build_1326 []byte, dm_build_1327 int, dm_build_1328 []byte, dm_build_1329 int, dm_build_1330 int) int {
	dm_build_1327 += Dm_build_1325.Dm_build_1307(dm_build_1326, dm_build_1327, uint32(dm_build_1330))
	return 4 + Dm_build_1325.Dm_build_1317(dm_build_1326, dm_build_1327, dm_build_1328, dm_build_1329, dm_build_1330)
}

func (Dm_build_1332 *dm_build_1260) Dm_build_1331(dm_build_1333 []byte, dm_build_1334 int, dm_build_1335 []byte, dm_build_1336 int, dm_build_1337 int) int {
	dm_build_1334 += Dm_build_1332.Dm_build_1302(dm_build_1333, dm_build_1334, uint16(dm_build_1337))
	return 2 + Dm_build_1332.Dm_build_1317(dm_build_1333, dm_build_1334, dm_build_1335, dm_build_1336, dm_build_1337)
}

func (Dm_build_1339 *dm_build_1260) Dm_build_1338(dm_build_1340 []byte, dm_build_1341 int, dm_build_1342 string, dm_build_1343 string, dm_build_1344 *DmConnection) int {
	dm_build_1345 := Dm_build_1339.Dm_build_1471(dm_build_1342, dm_build_1343, dm_build_1344)
	dm_build_1341 += Dm_build_1339.Dm_build_1307(dm_build_1340, dm_build_1341, uint32(len(dm_build_1345)))
	return 4 + Dm_build_1339.Dm_build_1317(dm_build_1340, dm_build_1341, dm_build_1345, 0, len(dm_build_1345))
}

func (Dm_build_1347 *dm_build_1260) Dm_build_1346(dm_build_1348 []byte, dm_build_1349 int, dm_build_1350 string, dm_build_1351 string, dm_build_1352 *DmConnection) int {
	dm_build_1353 := Dm_build_1347.Dm_build_1471(dm_build_1350, dm_build_1351, dm_build_1352)

	dm_build_1349 += Dm_build_1347.Dm_build_1302(dm_build_1348, dm_build_1349, uint16(len(dm_build_1353)))
	return 2 + Dm_build_1347.Dm_build_1317(dm_build_1348, dm_build_1349, dm_build_1353, 0, len(dm_build_1353))
}

func (Dm_build_1355 *dm_build_1260) Dm_build_1354(dm_build_1356 []byte, dm_build_1357 int) byte {
	return dm_build_1356[dm_build_1357]
}

func (Dm_build_1359 *dm_build_1260) Dm_build_1358(dm_build_1360 []byte, dm_build_1361 int) int16 {
	var dm_build_1362 int16
	dm_build_1362 = int16(dm_build_1360[dm_build_1361] & 0xff)
	dm_build_1361++
	dm_build_1362 |= int16(dm_build_1360[dm_build_1361]&0xff) << 8
	return dm_build_1362
}

func (Dm_build_1364 *dm_build_1260) Dm_build_1363(dm_build_1365 []byte, dm_build_1366 int) int32 {
	var dm_build_1367 int32
	dm_build_1367 = int32(dm_build_1365[dm_build_1366] & 0xff)
	dm_build_1366++
	dm_build_1367 |= int32(dm_build_1365[dm_build_1366]&0xff) << 8
	dm_build_1366++
	dm_build_1367 |= int32(dm_build_1365[dm_build_1366]&0xff) << 16
	dm_build_1366++
	dm_build_1367 |= int32(dm_build_1365[dm_build_1366]&0xff) << 24
	return dm_build_1367
}

func (Dm_build_1369 *dm_build_1260) Dm_build_1368(dm_build_1370 []byte, dm_build_1371 int) int64 {
	var dm_build_1372 int64
	dm_build_1372 = int64(dm_build_1370[dm_build_1371] & 0xff)
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 8
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 16
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 24
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 32
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 40
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 48
	dm_build_1371++
	dm_build_1372 |= int64(dm_build_1370[dm_build_1371]&0xff) << 56
	return dm_build_1372
}

func (Dm_build_1374 *dm_build_1260) Dm_build_1373(dm_build_1375 []byte, dm_build_1376 int) float32 {
	return math.Float32frombits(Dm_build_1374.Dm_build_1390(dm_build_1375, dm_build_1376))
}

func (Dm_build_1378 *dm_build_1260) Dm_build_1377(dm_build_1379 []byte, dm_build_1380 int) float64 {
	return math.Float64frombits(Dm_build_1378.Dm_build_1395(dm_build_1379, dm_build_1380))
}

func (Dm_build_1382 *dm_build_1260) Dm_build_1381(dm_build_1383 []byte, dm_build_1384 int) uint8 {
	return uint8(dm_build_1383[dm_build_1384] & 0xff)
}

func (Dm_build_1386 *dm_build_1260) Dm_build_1385(dm_build_1387 []byte, dm_build_1388 int) uint16 {
	var dm_build_1389 uint16
	dm_build_1389 = uint16(dm_build_1387[dm_build_1388] & 0xff)
	dm_build_1388++
	dm_build_1389 |= uint16(dm_build_1387[dm_build_1388]&0xff) << 8
	return dm_build_1389
}

func (Dm_build_1391 *dm_build_1260) Dm_build_1390(dm_build_1392 []byte, dm_build_1393 int) uint32 {
	var dm_build_1394 uint32
	dm_build_1394 = uint32(dm_build_1392[dm_build_1393] & 0xff)
	dm_build_1393++
	dm_build_1394 |= uint32(dm_build_1392[dm_build_1393]&0xff) << 8
	dm_build_1393++
	dm_build_1394 |= uint32(dm_build_1392[dm_build_1393]&0xff) << 16
	dm_build_1393++
	dm_build_1394 |= uint32(dm_build_1392[dm_build_1393]&0xff) << 24
	return dm_build_1394
}

func (Dm_build_1396 *dm_build_1260) Dm_build_1395(dm_build_1397 []byte, dm_build_1398 int) uint64 {
	var dm_build_1399 uint64
	dm_build_1399 = uint64(dm_build_1397[dm_build_1398] & 0xff)
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 8
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 16
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 24
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 32
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 40
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 48
	dm_build_1398++
	dm_build_1399 |= uint64(dm_build_1397[dm_build_1398]&0xff) << 56
	return dm_build_1399
}

func (Dm_build_1401 *dm_build_1260) Dm_build_1400(dm_build_1402 []byte, dm_build_1403 int) []byte {
	dm_build_1404 := Dm_build_1401.Dm_build_1390(dm_build_1402, dm_build_1403)
	return dm_build_1402[dm_build_1403+4 : dm_build_1403+4+int(dm_build_1404)]

}

func (Dm_build_1406 *dm_build_1260) Dm_build_1405(dm_build_1407 []byte, dm_build_1408 int) []byte {
	dm_build_1409 := Dm_build_1406.Dm_build_1385(dm_build_1407, dm_build_1408)
	return dm_build_1407[dm_build_1408+2 : dm_build_1408+2+int(dm_build_1409)]

}

func (Dm_build_1411 *dm_build_1260) Dm_build_1410(dm_build_1412 []byte, dm_build_1413 int, dm_build_1414 int) []byte {
	return dm_build_1412[dm_build_1413 : dm_build_1413+dm_build_1414]

}

func (Dm_build_1416 *dm_build_1260) Dm_build_1415(dm_build_1417 []byte, dm_build_1418 int, dm_build_1419 int, dm_build_1420 string, dm_build_1421 *DmConnection) string {
	return Dm_build_1416.Dm_build_1508(dm_build_1417[dm_build_1418:dm_build_1418+dm_build_1419], dm_build_1420, dm_build_1421)
}

func (Dm_build_1423 *dm_build_1260) Dm_build_1422(dm_build_1424 []byte, dm_build_1425 int, dm_build_1426 string, dm_build_1427 *DmConnection) string {
	dm_build_1428 := Dm_build_1423.Dm_build_1390(dm_build_1424, dm_build_1425)
	dm_build_1425 += 4
	return Dm_build_1423.Dm_build_1415(dm_build_1424, dm_build_1425, int(dm_build_1428), dm_build_1426, dm_build_1427)
}

func (Dm_build_1430 *dm_build_1260) Dm_build_1429(dm_build_1431 []byte, dm_build_1432 int, dm_build_1433 string, dm_build_1434 *DmConnection) string {
	dm_build_1435 := Dm_build_1430.Dm_build_1385(dm_build_1431, dm_build_1432)
	dm_build_1432 += 2
	return Dm_build_1430.Dm_build_1415(dm_build_1431, dm_build_1432, int(dm_build_1435), dm_build_1433, dm_build_1434)
}

func (Dm_build_1437 *dm_build_1260) Dm_build_1436(dm_build_1438 byte) []byte {
	return []byte{dm_build_1438}
}

func (Dm_build_1440 *dm_build_1260) Dm_build_1439(dm_build_1441 int16) []byte {
	return []byte{byte(dm_build_1441), byte(dm_build_1441 >> 8)}
}

func (Dm_build_1443 *dm_build_1260) Dm_build_1442(dm_build_1444 int32) []byte {
	return []byte{byte(dm_build_1444), byte(dm_build_1444 >> 8), byte(dm_build_1444 >> 16), byte(dm_build_1444 >> 24)}
}

func (Dm_build_1446 *dm_build_1260) Dm_build_1445(dm_build_1447 int64) []byte {
	return []byte{byte(dm_build_1447), byte(dm_build_1447 >> 8), byte(dm_build_1447 >> 16), byte(dm_build_1447 >> 24), byte(dm_build_1447 >> 32),
		byte(dm_build_1447 >> 40), byte(dm_build_1447 >> 48), byte(dm_build_1447 >> 56)}
}

func (Dm_build_1449 *dm_build_1260) Dm_build_1448(dm_build_1450 float32) []byte {
	return Dm_build_1449.Dm_build_1460(math.Float32bits(dm_build_1450))
}

func (Dm_build_1452 *dm_build_1260) Dm_build_1451(dm_build_1453 float64) []byte {
	return Dm_build_1452.Dm_build_1463(math.Float64bits(dm_build_1453))
}

func (Dm_build_1455 *dm_build_1260) Dm_build_1454(dm_build_1456 uint8) []byte {
	return []byte{byte(dm_build_1456)}
}

func (Dm_build_1458 *dm_build_1260) Dm_build_1457(dm_build_1459 uint16) []byte {
	return []byte{byte(dm_build_1459), byte(dm_build_1459 >> 8)}
}

func (Dm_build_1461 *dm_build_1260) Dm_build_1460(dm_build_1462 uint32) []byte {
	return []byte{byte(dm_build_1462), byte(dm_build_1462 >> 8), byte(dm_build_1462 >> 16), byte(dm_build_1462 >> 24)}
}

func (Dm_build_1464 *dm_build_1260) Dm_build_1463(dm_build_1465 uint64) []byte {
	return []byte{byte(dm_build_1465), byte(dm_build_1465 >> 8), byte(dm_build_1465 >> 16), byte(dm_build_1465 >> 24), byte(dm_build_1465 >> 32), byte(dm_build_1465 >> 40), byte(dm_build_1465 >> 48), byte(dm_build_1465 >> 56)}
}

func (Dm_build_1467 *dm_build_1260) Dm_build_1466(dm_build_1468 []byte, dm_build_1469 string, dm_build_1470 *DmConnection) []byte {
	if dm_build_1469 == "UTF-8" {
		return dm_build_1468
	}

	if dm_build_1470 == nil {
		if e := dm_build_1513(dm_build_1469); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1468), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1470.encodeBuffer == nil {
		dm_build_1470.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1470.encode = dm_build_1513(dm_build_1470.getServerEncoding())
		dm_build_1470.transformReaderDst = make([]byte, 4096)
		dm_build_1470.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1470.encode; e != nil {

		dm_build_1470.encodeBuffer.Reset()

		n, err := dm_build_1470.encodeBuffer.ReadFrom(
			Dm_build_1527(bytes.NewReader(dm_build_1468), e.NewEncoder(), dm_build_1470.transformReaderDst, dm_build_1470.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1470.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1472 *dm_build_1260) Dm_build_1471(dm_build_1473 string, dm_build_1474 string, dm_build_1475 *DmConnection) []byte {
	return Dm_build_1472.Dm_build_1466([]byte(dm_build_1473), dm_build_1474, dm_build_1475)
}

func (Dm_build_1477 *dm_build_1260) Dm_build_1476(dm_build_1478 []byte) byte {
	return Dm_build_1477.Dm_build_1354(dm_build_1478, 0)
}

func (Dm_build_1480 *dm_build_1260) Dm_build_1479(dm_build_1481 []byte) int16 {
	return Dm_build_1480.Dm_build_1358(dm_build_1481, 0)
}

func (Dm_build_1483 *dm_build_1260) Dm_build_1482(dm_build_1484 []byte) int32 {
	return Dm_build_1483.Dm_build_1363(dm_build_1484, 0)
}

func (Dm_build_1486 *dm_build_1260) Dm_build_1485(dm_build_1487 []byte) int64 {
	return Dm_build_1486.Dm_build_1368(dm_build_1487, 0)
}

func (Dm_build_1489 *dm_build_1260) Dm_build_1488(dm_build_1490 []byte) float32 {
	return Dm_build_1489.Dm_build_1373(dm_build_1490, 0)
}

func (Dm_build_1492 *dm_build_1260) Dm_build_1491(dm_build_1493 []byte) float64 {
	return Dm_build_1492.Dm_build_1377(dm_build_1493, 0)
}

func (Dm_build_1495 *dm_build_1260) Dm_build_1494(dm_build_1496 []byte) uint8 {
	return Dm_build_1495.Dm_build_1381(dm_build_1496, 0)
}

func (Dm_build_1498 *dm_build_1260) Dm_build_1497(dm_build_1499 []byte) uint16 {
	return Dm_build_1498.Dm_build_1385(dm_build_1499, 0)
}

func (Dm_build_1501 *dm_build_1260) Dm_build_1500(dm_build_1502 []byte) uint32 {
	return Dm_build_1501.Dm_build_1390(dm_build_1502, 0)
}

func (Dm_build_1504 *dm_build_1260) Dm_build_1503(dm_build_1505 []byte, dm_build_1506 string, dm_build_1507 *DmConnection) []byte {
	if dm_build_1506 == "UTF-8" {
		return dm_build_1505
	}

	if dm_build_1507 == nil {
		if e := dm_build_1513(dm_build_1506); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1505), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1507.encodeBuffer == nil {
		dm_build_1507.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1507.encode = dm_build_1513(dm_build_1507.getServerEncoding())
		dm_build_1507.transformReaderDst = make([]byte, 4096)
		dm_build_1507.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1507.encode; e != nil {

		dm_build_1507.encodeBuffer.Reset()

		n, err := dm_build_1507.encodeBuffer.ReadFrom(
			Dm_build_1527(bytes.NewReader(dm_build_1505), e.NewDecoder(), dm_build_1507.transformReaderDst, dm_build_1507.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1507.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1509 *dm_build_1260) Dm_build_1508(dm_build_1510 []byte, dm_build_1511 string, dm_build_1512 *DmConnection) string {
	return string(Dm_build_1509.Dm_build_1503(dm_build_1510, dm_build_1511, dm_build_1512))
}

func dm_build_1513(dm_build_1514 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1514); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1515 struct {
	dm_build_1516 io.Reader
	dm_build_1517 transform.Transformer
	dm_build_1518 error

	dm_build_1519                []byte
	dm_build_1520, dm_build_1521 int

	dm_build_1522                []byte
	dm_build_1523, dm_build_1524 int

	dm_build_1525 bool
}

const dm_build_1526 = 4096

func Dm_build_1527(dm_build_1528 io.Reader, dm_build_1529 transform.Transformer, dm_build_1530 []byte, dm_build_1531 []byte) *Dm_build_1515 {
	dm_build_1529.Reset()
	return &Dm_build_1515{
		dm_build_1516: dm_build_1528,
		dm_build_1517: dm_build_1529,
		dm_build_1519: dm_build_1530,
		dm_build_1522: dm_build_1531,
	}
}

func (dm_build_1533 *Dm_build_1515) Read(dm_build_1534 []byte) (int, error) {
	dm_build_1535, dm_build_1536 := 0, error(nil)
	for {

		if dm_build_1533.dm_build_1520 != dm_build_1533.dm_build_1521 {
			dm_build_1535 = copy(dm_build_1534, dm_build_1533.dm_build_1519[dm_build_1533.dm_build_1520:dm_build_1533.dm_build_1521])
			dm_build_1533.dm_build_1520 += dm_build_1535
			if dm_build_1533.dm_build_1520 == dm_build_1533.dm_build_1521 && dm_build_1533.dm_build_1525 {
				return dm_build_1535, dm_build_1533.dm_build_1518
			}
			return dm_build_1535, nil
		} else if dm_build_1533.dm_build_1525 {
			return 0, dm_build_1533.dm_build_1518
		}

		if dm_build_1533.dm_build_1523 != dm_build_1533.dm_build_1524 || dm_build_1533.dm_build_1518 != nil {
			dm_build_1533.dm_build_1520 = 0
			dm_build_1533.dm_build_1521, dm_build_1535, dm_build_1536 = dm_build_1533.dm_build_1517.Transform(dm_build_1533.dm_build_1519, dm_build_1533.dm_build_1522[dm_build_1533.dm_build_1523:dm_build_1533.dm_build_1524], dm_build_1533.dm_build_1518 == io.EOF)
			dm_build_1533.dm_build_1523 += dm_build_1535

			switch {
			case dm_build_1536 == nil:
				if dm_build_1533.dm_build_1523 != dm_build_1533.dm_build_1524 {
					dm_build_1533.dm_build_1518 = nil
				}

				dm_build_1533.dm_build_1525 = dm_build_1533.dm_build_1518 != nil
				continue
			case dm_build_1536 == transform.ErrShortDst && (dm_build_1533.dm_build_1521 != 0 || dm_build_1535 != 0):

				continue
			case dm_build_1536 == transform.ErrShortSrc && dm_build_1533.dm_build_1524-dm_build_1533.dm_build_1523 != len(dm_build_1533.dm_build_1522) && dm_build_1533.dm_build_1518 == nil:

			default:
				dm_build_1533.dm_build_1525 = true

				if dm_build_1533.dm_build_1518 == nil || dm_build_1533.dm_build_1518 == io.EOF {
					dm_build_1533.dm_build_1518 = dm_build_1536
				}
				continue
			}
		}

		if dm_build_1533.dm_build_1523 != 0 {
			dm_build_1533.dm_build_1523, dm_build_1533.dm_build_1524 = 0, copy(dm_build_1533.dm_build_1522, dm_build_1533.dm_build_1522[dm_build_1533.dm_build_1523:dm_build_1533.dm_build_1524])
		}
		dm_build_1535, dm_build_1533.dm_build_1518 = dm_build_1533.dm_build_1516.Read(dm_build_1533.dm_build_1522[dm_build_1533.dm_build_1524:])
		dm_build_1533.dm_build_1524 += dm_build_1535
	}
}
