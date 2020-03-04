package job

import (
	"github.com/hailongz/golang/db"
)

type Log struct {
	db.Object
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID" index:"ASC"`
	Appid	int64	`json:"appid" name:"appid" title:"应用ID" index:"ASC"`
	Sid	int64	`json:"sid" name:"sid" title:"主机ID" index:"ASC"`
	Type	int32	`json:"type" name:"type" title:"类型" index:"ASC"`
	Body	string	`json:"body" name:"body" title:"日志内容" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *Log) GetName() string {
	return "log"
}

func (O *Log) GetTitle() string {
	return "日志"
}

