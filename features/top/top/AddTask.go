package top

type AddTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Rate	int32	`json:"rate" name:"rate" title:"权重"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"搜索关键字"`
	Time	interface{}	`json:"time,omitempty" name:"time" title:"时间默认当前时间"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *AddTask) GetName() string {
	return "add.json"
}

func (T *AddTask) GetTitle() string {
	return "添加"
}

