package job

type SlaveRmTask struct {
	Id	int64	`json:"id" name:"id" title:"主机ID"`
}

func (T *SlaveRmTask) GetName() string {
	return "slave/rm.json"
}

func (T *SlaveRmTask) GetTitle() string {
	return "删除主机"
}

