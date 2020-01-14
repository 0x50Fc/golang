package inbox

type QueryTask struct {
	Uid     int64       `json:"uid" name:"uid" title:"用户ID"`
	Fuid    interface{} `json:"fuid,omitempty" name:"fuid" title:"发布者ID"`
	Type    interface{} `json:"type,omitempty" name:"type" title:"类型 type1 | type2 | type3"`
	Mid     interface{} `json:"mid,omitempty" name:"mid" title:"内容ID"`
	Iid     interface{} `json:"iid,omitempty" name:"iid" title:"内容项ID"`
	P       interface{} `json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N       interface{} `json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
	TopId   interface{} `json:"topId,omitempty" name:"topid" title:"顶部ID"`
	GroupBy interface{} `json:"groupBy,omitempty" name:"groupby" title:"分组"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}
