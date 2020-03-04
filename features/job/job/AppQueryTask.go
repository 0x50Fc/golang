package job

type AppQueryTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *AppQueryTask) GetName() string {
	return "app/query.json"
}

func (T *AppQueryTask) GetTitle() string {
	return "查询应用"
}

