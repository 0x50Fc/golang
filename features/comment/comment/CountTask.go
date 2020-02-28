package comment

type CountTask struct {
	Id   interface{} `json:"id,omitempty" name:"id" title:"评论ID"`
	Ids  interface{} `json:"ids,omitempty" name:"ids" title:"评论ID 多个逗号分割"`
	Pid  interface{} `json:"pid,omitempty" name:"pid" title:"父级别"`
	Eid  int64       `json:"eid" name:"eid" title:"评论目标ID"`
	Uid  interface{} `json:"uid,omitempty" name:"uid" title:"用户ID,0不验证"`
	Q    interface{} `json:"q,omitempty" name:"q" title:"内容模糊查询"`
	Path interface{} `json:"path,omitempty" name:"path" title:"path模糊查询,查询一个评论下的所有回复"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "获取评论数量"
}
