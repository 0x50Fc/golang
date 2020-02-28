package media

type GetTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"存储表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"存储分区"`
	Id	int64	`json:"id" name:"id" title:"媒体ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取媒体"
}

