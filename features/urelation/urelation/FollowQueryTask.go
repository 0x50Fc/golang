package urelation

type FollowQueryTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	In	interface{}	`json:"in,omitempty" title:"好友ID，多个逗号分割"`
	Q	interface{}	`json:"q,omitempty" title:"模糊匹配关键字"`
	P	interface{}	`json:"p,omitempty" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" title:"分页大小，默认 20"`
}

func (T *FollowQueryTask) GetName() string {
	return "follow/query.json"
}

func (T *FollowQueryTask) GetTitle() string {
	return "查询关注的好友"
}

