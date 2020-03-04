package job

type SlaveCountTask struct {
}

func (T *SlaveCountTask) GetName() string {
	return "slave/count.json"
}

func (T *SlaveCountTask) GetTitle() string {
	return "主机数量"
}

