package member

type QueryTask struct {
	Bid	interface{}	`json:"bid,omitempty" name:"bid" title:"\b\b商户ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"成员ID"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

