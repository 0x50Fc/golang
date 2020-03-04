package wx

type QueryTask struct {
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型,多个逗号分割"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"appid"`
	Openid	interface{}	`json:"openid,omitempty" name:"openid" title:"openid"`
	Unionid	interface{}	`json:"unionid,omitempty" name:"unionid" title:"unionid"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"状态 多个逗号分割"`
	Bind	interface{}	`json:"bind,omitempty" name:"bind" title:"是否绑定"`
	StartTime	interface{}	`json:"startTime,omitempty" name:"starttime" title:"绑定开始时间"`
	EndTime	interface{}	`json:"endTime,omitempty" name:"endtime" title:"绑定结束时间"`
	Info	interface{}	`json:"info,omitempty" name:"info" title:"是否有用户信息"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

