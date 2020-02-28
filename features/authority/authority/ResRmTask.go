package authority

type ResRmTask struct {
	Id	int64	`json:"id" name:"id" title:"资源ID"`
}

func (T *ResRmTask) GetName() string {
	return "res/rm.json"
}

func (T *ResRmTask) GetTitle() string {
	return "删除资源"
}

