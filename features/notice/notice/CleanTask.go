package notice

type CleanTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型, 多个逗号分割"`
	Fid	interface{}	`json:"fid,omitempty" name:"fid" title:"消息来源ID , 多个逗号分割"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"消息来源项ID , 多个逗号分割"`
}

func (T *CleanTask) GetName() string {
	return "clean.json"
}

func (T *CleanTask) GetTitle() string {
	return "清理消息"
}

