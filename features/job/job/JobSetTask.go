package job

type JobSetTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"应用ID"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	MaxCount	interface{}	`json:"maxCount,omitempty" name:"maxcount" title:"总任务数"`
	Count	interface{}	`json:"count,omitempty" name:"count" title:"已执行任务数"`
	ErrCount	interface{}	`json:"errCount,omitempty" name:"errcount" title:"错误任务数"`
	AddCount	interface{}	`json:"addCount,omitempty" name:"addcount" title:"增加已执行任务数"`
	AddErrCount	interface{}	`json:"addErrCount,omitempty" name:"adderrcount" title:"增加错误任务数"`
	Stime	interface{}	`json:"stime,omitempty" name:"stime" title:"开始时间"`
}

func (T *JobSetTask) GetName() string {
	return "job/set.json"
}

func (T *JobSetTask) GetTitle() string {
	return "修改工作"
}

