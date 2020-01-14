package inbox

type AddTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	int64	`json:"type" name:"type" title:"类型"`
	Mid	int64	`json:"mid" name:"mid" title:"内容ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"内容项ID"`
	Fuid	int64	`json:"fuid" name:"fuid" title:"发布者ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	Ctime	interface{}	`json:"ctime,omitempty" name:"ctime" title:"创建时间"`
}

func (T *AddTask) GetName() string {
	return "add.json"
}

func (T *AddTask) GetTitle() string {
	return "添加"
}

