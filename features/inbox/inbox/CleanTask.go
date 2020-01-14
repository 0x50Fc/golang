package inbox

type CleanTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型 type1 | type2 | type3"`
	Mid	interface{}	`json:"mid,omitempty" name:"mid" title:"内容ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"发布者ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"内容项ID"`
}

func (T *CleanTask) GetName() string {
	return "clean.json"
}

func (T *CleanTask) GetTitle() string {
	return "清理"
}

