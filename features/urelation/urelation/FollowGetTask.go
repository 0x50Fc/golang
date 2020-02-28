package urelation

type FollowGetTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	Fuid	int64	`json:"fuid,omitempty" title:"好友ID"`
}

func (T *FollowGetTask) GetName() string {
	return "follow/get.json"
}

func (T *FollowGetTask) GetTitle() string {
	return "获取关注的好友"
}

