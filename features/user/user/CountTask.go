package user

type CountTask struct {
	Ids	interface{}	`json:"ids,omitempty" name:"ids" title:"用户ID,逗号分割"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"用户名"`
	Nick	interface{}	`json:"nick,omitempty" name:"nick" title:"昵称"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"用户名前缀"`
	Suffix	interface{}	`json:"suffix,omitempty" name:"suffix" title:"用户名后缀"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "用户数量"
}

