package media

type QueryTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"存储表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"存储分区"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"路径前缀"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型,多个逗号分割"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

