package notice

type CreateTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"通知类型 默认 0"`
	Fid	interface{}	`json:"fid,omitempty" name:"fid" title:"消息来源ID 默认 0"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"消息来源项ID 默认 0"`
	Body	string	`json:"body" name:"body" title:"通知内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建"
}

