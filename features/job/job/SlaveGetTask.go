package job

type SlaveGetTask struct {
	Id	int64	`json:"id" name:"id" title:"主机ID"`
}

func (T *SlaveGetTask) GetName() string {
	return "slave/get.json"
}

func (T *SlaveGetTask) GetTitle() string {
	return "获取主机"
}

