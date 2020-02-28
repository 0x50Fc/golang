package addup

type UpTask struct {
	Name	string	`json:"name" name:"name" title:"统计表名"`
	Region	interface{}	`json:"region,omitempty" name:"region" title:"分区 默认无分区"`
	Iid	int64	`json:"iid" name:"iid" title:"统计项ID"`
	UnionKeys	interface{}	`json:"unionKeys,omitempty" name:"unionkeys" title:"iid + time + unionKeys唯一\nJSON 对象"`
	Set	interface{}	`json:"set,omitempty" name:"set" title:"设置数据项的值 JSON\n{ \"key\" : \"value\" }"`
	Add	interface{}	`json:"add,omitempty" name:"add" title:"增加数据项的值 JSON\n{ \"key\" : \"value\" }"`
	Time	int64	`json:"time" name:"time" title:"时间 (秒)"`
}

func (T *UpTask) GetName() string {
	return "up.json"
}

func (T *UpTask) GetTitle() string {
	return "上传统计数据"
}

