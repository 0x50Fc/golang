package job

import (
	"github.com/hailongz/golang/db"
)

type Job struct {
	db.Object
	Alias	string	`json:"alias" name:"alias" title:"别名" length:"128" index:"ASC"`
	Type	int32	`json:"type" name:"type" title:"类型" index:"ASC"`
	State	int32	`json:"state" name:"state" title:"状态" index:"ASC"`
	Appid	int64	`json:"appid" name:"appid" title:"应用ID" index:"ASC"`
	Sid	int64	`json:"sid" name:"sid" title:"主机ID" index:"ASC"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	MaxCount	int32	`json:"maxCount" name:"maxcount" title:"总任务数"`
	Count	int32	`json:"count" name:"count" title:"已执行任务数"`
	ErrCount	int32	`json:"errCount" name:"errcount" title:"错误任务数"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
	Stime	int64	`json:"stime" name:"stime" title:"开始时间" index:"ASC"`
}

func (O *Job) GetName() string {
	return "job"
}

func (O *Job) GetTitle() string {
	return "工作"
}

