package app

type VerSetTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
	Ver	int32	`json:"ver" name:"ver" title:"版本号"`
	Info	interface{}	`json:"info,omitempty" name:"info" title:"INFO"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据"`
}

func (T *VerSetTask) GetName() string {
	return "ver/set.json"
}

func (T *VerSetTask) GetTitle() string {
	return "删除版本"
}

