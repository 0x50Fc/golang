package job

type SlaveJobGetTask struct {
	Token	string	`json:"token" name:"token" title:"主机 token"`
	Expires	int64	`json:"expires" name:"expires" title:"主机超时时间 秒"`
}

func (T *SlaveJobGetTask) GetName() string {
	return "slave/job/get.json"
}

func (T *SlaveJobGetTask) GetTitle() string {
	return "主机获取可用工作"
}

