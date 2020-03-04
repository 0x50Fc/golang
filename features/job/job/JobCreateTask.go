package job

type JobCreateTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"应用ID"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	Stime	interface{}	`json:"stime,omitempty" name:"stime" title:"开始时间"`
}

func (T *JobCreateTask) GetName() string {
	return "job/create.json"
}

func (T *JobCreateTask) GetTitle() string {
	return "创建工作"
}

