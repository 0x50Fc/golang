package inbox

type CountTask struct {
	Uid     int64       `json:"uid" name:"uid" title:"用户ID"`
	Fuid    interface{} `json:"fuid,omitempty" name:"fuid" title:"发布者ID"`
	Type    interface{} `json:"type,omitempty" name:"type" title:"类型 type1 | type2 | type3"`
	Mid     interface{} `json:"mid,omitempty" name:"mid" title:"内容ID"`
	Iid     interface{} `json:"iid,omitempty" name:"iid" title:"内容项ID"`
	TopId   interface{} `json:"topId,omitempty" name:"topid" title:"顶部ID"`
	GroupBy interface{} `json:"groupBy,omitempty" name:"groupby" title:"分组"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}
