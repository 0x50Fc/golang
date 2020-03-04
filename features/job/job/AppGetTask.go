package job

type AppGetTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
}

func (T *AppGetTask) GetName() string {
	return "app/get.json"
}

func (T *AppGetTask) GetTitle() string {
	return "获取应用"
}

