package inbox

type UserQueryTask struct {
	Type int64       `json:"type" name:"type" title:"类型"`
	Mid  interface{} `json:"mid,omitempty" name:"mid" title:"内容ID"`
	Iid  interface{} `json:"iid,omitempty" name:"iid" title:"内容项ID"`
	P    interface{} `json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N    interface{} `json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *UserQueryTask) GetName() string {
	return "user/query.json"
}

func (T *UserQueryTask) GetTitle() string {
	return "查询订阅用户"
}
