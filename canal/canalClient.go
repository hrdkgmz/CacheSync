package canal

import (
	"fmt"
	"github.com/hrdkgmz/cacheSync/task"
	"github.com/hrdkgmz/cacheSync/taskHandle"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"
)

func StartFetchEvent() {

	// 192.168.199.17 替换成你的canal server的地址
	// example 替换成-e canal.destinations=example 你自己定义的名字
	connector := client.NewSimpleCanalConnector("192.168.110.153", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)

	err = connector.Subscribe("bctest\\..*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("===没有数据了===")
			continue
		}

		handleMsg(message.Entries)

	}
}

func handleMsg(entrys []protocol.Entry) {

	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		if err!=nil{
			return
		}
		if rowChange != nil {
			header := entry.GetHeader()
			handleRows(rowChange,header)
		}
	}
}

func handleRows(rowChange *protocol.RowChange, header *protocol.Header){
	eventType := rowChange.GetEventType()
	fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))
	taskHandler:= taskHandle.GetInstance()
	for _, rowData := range rowChange.GetRowDatas() {
		if eventType == protocol.EventType_DELETE {
			printColumn(rowData.GetBeforeColumns())
			taskHandler.Do(task.NewDeleteTask(rowData))
		} else if eventType == protocol.EventType_INSERT {
			printColumn(rowData.GetAfterColumns())
			taskHandler.Do(task.NewInsertTask(rowData))
		} else if eventType==protocol.EventType_UPDATE {
			fmt.Println("-------> before")
			printColumn(rowData.GetBeforeColumns())
			fmt.Println("-------> after")
			printColumn(rowData.GetAfterColumns())
			taskHandler.Do(task.NewUpdateTask(rowData))
		}else{
			log.Println("不支持的事件类型："+strconv.Itoa((int(eventType))))
		}
	}

}

func printColumn(columns []*protocol.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

//func checkError(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}