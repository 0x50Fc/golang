package media

type SetTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"存储表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"存储分区"`
	Id	int64	`json:"id" name:"id" title:"媒体ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"标题"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"关键字"`
	Path	interface{}	`json:"path,omitempty" name:"path" title:"存储路径"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

