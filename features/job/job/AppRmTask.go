package job

type AppRmTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
}

func (T *AppRmTask) GetName() string {
	return "app/rm.json"
}

func (T *AppRmTask) GetTitle() string {
	return "删除应用"
}

