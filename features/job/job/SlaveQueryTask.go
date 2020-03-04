package job

type SlaveQueryTask struct {
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *SlaveQueryTask) GetName() string {
	return "slave/query.json"
}

func (T *SlaveQueryTask) GetTitle() string {
	return "查询主机"
}

