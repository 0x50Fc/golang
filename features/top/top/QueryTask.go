package top

type QueryTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	Tids	interface{}	`json:"tids,omitempty" name:"tids" title:"目标ID,多个逗号分割"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
	TopId	interface{}	`json:"topId,omitempty" name:"topid" title:"顶部ID"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

