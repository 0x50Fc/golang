package article

type GetTask struct {
	Id	int64	`json:"id" name:"id" title:"动态ID"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取动态"
}

