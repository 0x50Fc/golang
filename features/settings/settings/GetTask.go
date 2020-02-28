package settings

type GetTask struct {
	Name	string	`json:"name" name:"name" title:"名称"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取"
}

