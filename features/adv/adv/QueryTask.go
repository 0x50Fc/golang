package adv

type QueryTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"广告ID"`
	Channel	interface{}	`json:"channel,omitempty" name:"channel" title:"频道"`
	Stime	interface{}	`json:"stime,omitempty" name:"stime" title:"开始时间"`
	Etime	interface{}	`json:"etime,omitempty" name:"etime" title:"结束时间"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "广告评论"
}

