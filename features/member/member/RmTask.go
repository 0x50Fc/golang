package member

type RmTask struct {
	Bid	int64	`json:"bid" name:"bid" title:"\b\b商户ID"`
	Uid	int64	`json:"uid" name:"uid" title:"成员ID"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除成员信息"
}

