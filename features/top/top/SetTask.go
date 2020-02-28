package top

type SetTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"搜索关键字"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

