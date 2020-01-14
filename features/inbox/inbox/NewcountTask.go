package inbox

type NewcountTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"发布者ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型 type1 | type2 | type3"`
	TopId	int64	`json:"topId" name:"topid" title:"顶部ID"`
	GroupBy	interface{}	`json:"groupBy,omitempty" name:"groupby" title:"分组"`
}

func (T *NewcountTask) GetName() string {
	return "newcount.json"
}

func (T *NewcountTask) GetTitle() string {
	return "最新数量"
}

