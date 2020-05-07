package lookin

type CodeTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Fcode	interface{}	`json:"fcode,omitempty" name:"fcode" title:"代码"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"好友ID"`
}

func (T *CodeTask) GetName() string {
	return "code.json"
}

func (T *CodeTask) GetTitle() string {
	return "生成推荐码"
}

