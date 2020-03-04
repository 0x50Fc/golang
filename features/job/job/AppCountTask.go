package job

type AppCountTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
}

func (T *AppCountTask) GetName() string {
	return "app/count.json"
}

func (T *AppCountTask) GetTitle() string {
	return "应用数量"
}

