package user

type QueryTask struct {
	Ids	interface{}	`json:"ids,omitempty" name:"ids" title:"用户ID,逗号分割"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"用户名"`
	Nick	interface{}	`json:"nick,omitempty" name:"nick" title:"昵称"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"用户名前缀"`
	Suffix	interface{}	`json:"suffix,omitempty" name:"suffix" title:"用户名后缀"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询用户"
}

