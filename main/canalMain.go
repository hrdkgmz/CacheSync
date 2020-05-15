package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/hrdkgmz/cacheSync/task"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"os"
	"strconv"
	"time"

	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
	"github.com/golang/protobuf/proto"
)

var (
	_ip          string = "192.168.110.164"
	_port        int    = 11111
	_username    string = ""
	_password    string = ""
	_destination string = "example"
	_soTimeOut   int32  = 60000
	_idleTimeOut int32  = 60 * 60 * 1000
	_subscribe   string = "bctest\\..*"
)

func StartFetchEvent() {
	connector := client.NewSimpleCanalConnector(_ip, _port, _username, _password, _destination, _soTimeOut, _idleTimeOut)
	err := connector.Connect()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	err = connector.Subscribe(_subscribe)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("Canal连接成功，开始监听binlog...")
	for {

		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Error(err)
			continue
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			//fmt.Println("===没有数据了===")
			continue
		}

		handleMsg(message.Entries)

	}
}

func handleMsg(entries []protocol.Entry) {

	for _, entry := range entries {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		if err != nil {
			return
		}
		if rowChange != nil {
			header := entry.GetHeader()
			handleRows(rowChange, header)
		}
	}
}

func handleRows(rowChange *protocol.RowChange, header *protocol.Header) {
	eventType := rowChange.GetEventType()
	log.Info(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))
	taskHandler := taskHandle.GetInstance()
	for _, rowData := range rowChange.GetRowDatas() {
		if eventType == protocol.EventType_DELETE {
			printColumn(rowData.GetBeforeColumns())
			taskHandler.Do(task.NewDeleteTask(rowData, header))
		} else if eventType == protocol.EventType_INSERT {
			printColumn(rowData.GetAfterColumns())
			taskHandler.Do(task.NewInsertTask(rowData, header))
		} else if eventType == protocol.EventType_UPDATE {
			log.Info("-------> before")
			printColumn(rowData.GetBeforeColumns())
			log.Info("-------> after")
			printColumn(rowData.GetAfterColumns())
			taskHandler.Do(task.NewUpdateTask(rowData, header))
		} else {
			log.Error("不支持的事件类型：" + strconv.Itoa(int(eventType)))
		}
	}

}

func printColumn(columns []*protocol.Column) {
	for _, col := range columns {
		log.Info(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

func SetCanalParameters(ip string, port int, user string, pass string, dest string, soTime int32, idleTime int32, schema string) {
	_ip = ip
	_port = port
	_password = pass
	_destination = dest
	_soTimeOut = soTime
	_idleTimeOut = idleTime
	_subscribe = schema + "\\..*"
}
