package comment

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"评论ID"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID,0不验证"`
	Body	interface{}	`json:"body,omitempty" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改评论"
}

