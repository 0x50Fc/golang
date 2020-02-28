package urelation

type CountTask struct {
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "获取数量"
}

