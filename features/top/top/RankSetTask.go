package top

type RankSetTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Rank	int32	`json:"rank" name:"rank" title:"排名 0 表示不指定排名"`
}

func (T *RankSetTask) GetName() string {
	return "rank/set.json"
}

func (T *RankSetTask) GetTitle() string {
	return "修改排名"
}

