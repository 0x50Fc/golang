package urelation

type FollowTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	Fuid	int64	`json:"fuid,omitempty" title:"好友ID"`
	Title	string	`json:"title,omitempty" title:"备注名"`
}

func (T *FollowTask) GetName() string {
	return "follow.json"
}

func (T *FollowTask) GetTitle() string {
	return "关注好友"
}

