package app

type VerRmTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
	Ver	int32	`json:"ver" name:"ver" title:"版本号"`
}

func (T *VerRmTask) GetName() string {
	return "ver/rm.json"
}

func (T *VerRmTask) GetTitle() string {
	return "删除版本"
}

