package job

type LogAddTask struct {
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID"`
	Appid	int64	`json:"appid" name:"appid" title:"应用ID"`
	Sid	int64	`json:"sid" name:"sid" title:"主机ID"`
	Type	int32	`json:"type" name:"type" title:"类型"`
	Body	string	`json:"body" name:"body" title:"日志内容"`
}

func (T *LogAddTask) GetName() string {
	return "log/add.json"
}

func (T *LogAddTask) GetTitle() string {
	return "添加日志"
}

