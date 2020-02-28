package top

type RankCountTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	Tids	interface{}	`json:"tids,omitempty" name:"tids" title:"目标ID,多个逗号分割"`
	TopId	interface{}	`json:"topId,omitempty" name:"topid" title:"顶部ID"`
}

func (T *RankCountTask) GetName() string {
	return "rank/count.json"
}

func (T *RankCountTask) GetTitle() string {
	return "排名数量"
}

