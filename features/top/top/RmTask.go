package top

type RmTask struct {
	Name	string	`json:"name" name:"name" title:"推荐表名"`
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

