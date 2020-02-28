package top

type GetTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Tid	int64	`json:"tid" name:"tid" title:"ID"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取推荐项"
}

