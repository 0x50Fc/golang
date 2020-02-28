package notice

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"分组ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"通知类型"`
	Fid	interface{}	`json:"fid,omitempty" name:"fid" title:"消息来源ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"消息来源项ID"`
	Body	interface{}	`json:"body,omitempty" name:"body" title:"通知内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

