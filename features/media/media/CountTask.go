package media

type CountTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"存储表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"存储分区"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"路径前缀"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型,多个逗号分割"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}

