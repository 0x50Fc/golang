package job

type SlaveJobUpTask struct {
	Token	string	`json:"token" name:"token" title:"主机 token"`
	Expires	int64	`json:"expires" name:"expires" title:"主机超时时间 秒"`
	JobId	int64	`json:"jobId" name:"jobid" title:"工作ID"`
	Done	bool	`json:"done" name:"done" title:"工作是否已完成"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	MaxCount	interface{}	`json:"maxCount,omitempty" name:"maxcount" title:"总任务数"`
	Count	interface{}	`json:"count,omitempty" name:"count" title:"已执行任务数"`
	ErrCount	interface{}	`json:"errCount,omitempty" name:"errcount" title:"错误任务数"`
	AddCount	interface{}	`json:"addCount,omitempty" name:"addcount" title:"增加已执行任务数"`
	AddErrCount	interface{}	`json:"addErrCount,omitempty" name:"adderrcount" title:"增加错误任务数"`
}

func (T *SlaveJobUpTask) GetName() string {
	return "slave/job/up.json"
}

func (T *SlaveJobUpTask) GetTitle() string {
	return "主机更新工作进度"
}

