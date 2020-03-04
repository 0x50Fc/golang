package job

type JobGetTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
}

func (T *JobGetTask) GetName() string {
	return "job/get.json"
}

func (T *JobGetTask) GetTitle() string {
	return "获取工作"
}

