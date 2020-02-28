package settings

type RmTask struct {
	Name	string	`json:"name" name:"name" title:"名称"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

