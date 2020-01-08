package app

type VerGetTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
	Ver	int32	`json:"ver" name:"ver" title:"版本号"`
}

func (T *VerGetTask) GetName() string {
	return "ver/get.json"
}

func (T *VerGetTask) GetTitle() string {
	return "获取版本"
}

