package job

type LogCountTask struct {
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"日志类型 多个都会分割"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"关键字"`
}

func (T *LogCountTask) GetName() string {
	return "log/count.json"
}

func (T *LogCountTask) GetTitle() string {
	return "日志数量"
}

