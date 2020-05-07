package member

type CountTask struct {
	Bid	interface{}	`json:"bid,omitempty" name:"bid" title:"\b\b商户ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"成员ID"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}

