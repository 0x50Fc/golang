package like

type QueryTask struct {
	Tid     int64       `json:"tid" name:"tid" title:"目标"`
	Iid     interface{} `json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid     interface{} `json:"uid,omitempty" name:"uid" title:"用户ID"`
	P       interface{} `json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N       interface{} `json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
	Maxtime interface{} `json:"maxtime,omitempty" name:"maxtime" title:"最大时间"`
	Mintime interface{} `json:"mintime,omitempty" name:"mintime" title:"最小时间"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}
