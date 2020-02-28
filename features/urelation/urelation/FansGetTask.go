package urelation

type FansGetTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	Fuid	int64	`json:"fuid,omitempty" title:"好友ID"`
}

func (T *FansGetTask) GetName() string {
	return "fans/get.json"
}

func (T *FansGetTask) GetTitle() string {
	return "获取粉丝"
}

