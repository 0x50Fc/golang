package top

type RankSortTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Limit	int32	`json:"limit" name:"limit" title:"限制数量"`
}

func (T *RankSortTask) GetName() string {
	return "rank/sort.json"
}

func (T *RankSortTask) GetTitle() string {
	return "计算排名"
}

