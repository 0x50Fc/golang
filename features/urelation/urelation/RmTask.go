package urelation

type RmTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	Fuid	int64	`json:"fuid,omitempty" title:"好友ID"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "取消关注"
}

