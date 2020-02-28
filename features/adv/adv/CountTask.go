package adv

type CountTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"广告ID"`
	Channel	interface{}	`json:"channel,omitempty" name:"channel" title:"频道"`
	Stime	interface{}	`json:"stime,omitempty" name:"stime" title:"开始时间"`
	Etime	interface{}	`json:"etime,omitempty" name:"etime" title:"结束时间"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "获取广告数量"
}

