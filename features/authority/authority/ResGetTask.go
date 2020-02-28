package authority

type ResGetTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"资源ID"`
	Path	interface{}	`json:"path,omitempty" name:"path" title:"资源路径"`
}

func (T *ResGetTask) GetName() string {
	return "res/get.json"
}

func (T *ResGetTask) GetTitle() string {
	return "获取资源"
}

