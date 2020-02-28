package authority

type RmTask struct {
	Key	string	`json:"key" title:"键值"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

