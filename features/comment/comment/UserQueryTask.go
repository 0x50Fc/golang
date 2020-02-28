package comment

type UserQueryTask struct {
	Eid     int64       `json:"eid" name:"eid" title:"评论目标ID"`
	P       interface{} `json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N       interface{} `json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
	Maxtime interface{} `json:"maxtime,omitempty" name:"maxtime" title:"最大时间"`
	Mintime interface{} `json:"mintime,omitempty" name:"mintime" title:"最小时间"`
}

func (T *UserQueryTask) GetName() string {
	return "user/query.json"
}

func (T *UserQueryTask) GetTitle() string {
	return "查询评论的用户"
}
