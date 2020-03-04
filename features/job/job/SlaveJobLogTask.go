package job

type SlaveJobLogTask struct {
	Token	string	`json:"token" name:"token" title:"主机 token"`
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID"`
	Appid	int64	`json:"appid" name:"appid" title:"应用ID"`
	Type	int32	`json:"type" name:"type" title:"类型"`
	Body	string	`json:"body" name:"body" title:"日志内容"`
}

func (T *SlaveJobLogTask) GetName() string {
	return "slave/job/log.json"
}

func (T *SlaveJobLogTask) GetTitle() string {
	return "主机更新工作日志"
}

