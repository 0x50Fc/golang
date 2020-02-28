package notice

type NewcountByTypesTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Ids	interface{}	`json:"ids,omitempty" name:"ids" title:"ID,多个逗号分割"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型, 多个逗号分割"`
	Fid	interface{}	`json:"fid,omitempty" name:"fid" title:"消息来源ID , 多个逗号分割"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"消息来源项ID , 多个逗号分割"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	TopId	int64	`json:"topId" name:"topid" title:"顶部ID"`
}

func (T *NewcountByTypesTask) GetName() string {
	return "newcountByTypes.json"
}

func (T *NewcountByTypesTask) GetTitle() string {
	return "最新数量按类型统计"
}

