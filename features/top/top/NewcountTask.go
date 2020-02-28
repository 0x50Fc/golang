package top

type NewcountTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	TopId	int64	`json:"topId" name:"topid" title:"顶部ID"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
}

func (T *NewcountTask) GetName() string {
	return "newcount.json"
}

func (T *NewcountTask) GetTitle() string {
	return "最新数量"
}

