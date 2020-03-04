package job

type LogQueryTask struct {
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"日志类型 多个都会分割"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"关键字"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *LogQueryTask) GetName() string {
	return "log/query.json"
}

func (T *LogQueryTask) GetTitle() string {
	return "查询日志"
}

