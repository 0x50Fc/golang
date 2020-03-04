package job

type JobCountTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"应用ID"`
	Sid	interface{}	`json:"sid,omitempty" name:"sid" title:"主机ID"`
}

func (T *JobCountTask) GetName() string {
	return "job/count.json"
}

func (T *JobCountTask) GetTitle() string {
	return "工作数量"
}

