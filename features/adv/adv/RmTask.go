package adv

type RmTask struct {
	Id	int64	`json:"id" name:"id" title:"评论ID"`
	Channel	string	`json:"channel" name:"channel" title:"频道"`
	Position	int32	`json:"position" name:"position" title:"广告组位置"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除广告"
}

