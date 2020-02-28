package authority

type ResCountTask struct {
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"前缀"`
}

func (T *ResCountTask) GetName() string {
	return "res/count.json"
}

func (T *ResCountTask) GetTitle() string {
	return "资源数量"
}

