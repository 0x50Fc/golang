package job

type JobQueryTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"应用ID"`
	Sid	interface{}	`json:"sid,omitempty" name:"sid" title:"主机ID"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *JobQueryTask) GetName() string {
	return "job/query.json"
}

func (T *JobQueryTask) GetTitle() string {
	return "查询工作"
}

