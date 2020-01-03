package app

type VerQueryTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *VerQueryTask) GetName() string {
	return "ver/query.json"
}

func (T *VerQueryTask) GetTitle() string {
	return "查询用户"
}

