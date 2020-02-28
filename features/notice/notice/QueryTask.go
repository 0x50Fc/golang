package notice

type QueryTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Ids	interface{}	`json:"ids,omitempty" name:"ids" title:"ID,多个逗号分割"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型, 多个逗号分割"`
	Fid	interface{}	`json:"fid,omitempty" name:"fid" title:"消息来源ID , 多个逗号分割"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"消息来源项ID , 多个逗号分割"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	TopId	interface{}	`json:"topId,omitempty" name:"topid" title:"顶部ID"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

