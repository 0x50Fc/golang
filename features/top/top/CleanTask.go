package top

type CleanTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Limit	interface{}	`json:"limit,omitempty" name:"limit" title:"保留最大数量"`
}

func (T *CleanTask) GetName() string {
	return "clean.json"
}

func (T *CleanTask) GetTitle() string {
	return "清理"
}

