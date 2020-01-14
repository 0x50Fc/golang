package inbox

type GetTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取"
}

