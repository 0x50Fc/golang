package top

type RankNewcountTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	TopId	int64	`json:"topId" name:"topid" title:"顶部ID"`
}

func (T *RankNewcountTask) GetName() string {
	return "rank/newcount.json"
}

func (T *RankNewcountTask) GetTitle() string {
	return "排名最新数量"
}

