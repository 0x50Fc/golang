package app

type VerCountTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
}

func (T *VerCountTask) GetName() string {
	return "ver/count.json"
}

func (T *VerCountTask) GetTitle() string {
	return "用户数量"
}

