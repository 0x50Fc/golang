package inbox

type UserCountTask struct {
	Type int64       `json:"type" name:"type" title:"类型"`
	Mid  interface{} `json:"mid,omitempty" name:"mid" title:"内容ID"`
	Iid  interface{} `json:"iid,omitempty" name:"iid" title:"内容项ID"`
}

func (T *UserCountTask) GetName() string {
	return "user/count.json"
}

func (T *UserCountTask) GetTitle() string {
	return "用户数量"
}
