package top

type BatchAddTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Items	interface{}	`json:"items,omitempty" name:"items" title:"其他数据 JSON\n[ {tid : 1, rate : 1, options:  {}, keyword:''} ]"`
}

func (T *BatchAddTask) GetName() string {
	return "batch/add.json"
}

func (T *BatchAddTask) GetTitle() string {
	return "批量添加"
}

