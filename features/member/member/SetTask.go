package member

type SetTask struct {
	Bid	int64	`json:"bid" name:"bid" title:"\b\b商户ID"`
	Uid	int64	`json:"uid" name:"uid" title:"成员ID"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"备注名"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"搜索关键字"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改成员信息"
}

