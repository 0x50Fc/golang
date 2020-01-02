package user

type GetTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"用户ID"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"用户名"`
	Nick	interface{}	`json:"nick,omitempty" name:"nick" title:"昵称"`
	Autocreate	interface{}	`json:"autocreate,omitempty" name:"autocreate" title:"是否自动创建, name 必须存在"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取用户"
}

