package job

type JobCancelTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
}

func (T *JobCancelTask) GetName() string {
	return "job/cancel.json"
}

func (T *JobCancelTask) GetTitle() string {
	return "取消工作"
}

