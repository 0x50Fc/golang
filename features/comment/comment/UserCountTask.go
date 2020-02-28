package comment

type UserCountTask struct {
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
}

func (T *UserCountTask) GetName() string {
	return "user/count.json"
}

func (T *UserCountTask) GetTitle() string {
	return "获取评论用户数量"
}

