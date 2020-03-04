package job

type JobRmTask struct {
	Id	int64	`json:"id" name:"id" title:"工作ID"`
}

func (T *JobRmTask) GetName() string {
	return "job/rm.json"
}

func (T *JobRmTask) GetTitle() string {
	return "删除工作"
}

