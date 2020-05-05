package lookin

type RmTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"用户ID"`
	Flevel	interface{}	`json:"flevel,omitempty" name:"flevel" title:"好友级别，多个逗号分割"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "取消赞"
}

