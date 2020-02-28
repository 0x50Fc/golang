package adv

type AddTask struct {
	Title	string	`json:"title" name:"title" title:"标题"`
	Channel	string	`json:"channel" name:"channel" title:"频道" index:"ASC"`
	Position	int32	`json:"position" name:"position" title:"广告组位置"`
	Pic	string	`json:"pic" name:"pic" title:"图片"`
	Description	string	`json:"description" name:"description" title:"描述"`
	Link	string	`json:"link" name:"link" title:"跳转链接"`
	Linktype	int32	`json:"linktype" name:"linktype" title:"跳转类型"`
	Sort	int32	`json:"sort" name:"sort" title:"排序"`
	Starttime	int64	`json:"starttime" name:"starttime" title:"开始时间"`
	Endtime	int64	`json:"endtime" name:"endtime" title:"结束时间"`
}

func (T *AddTask) GetName() string {
	return "add.json"
}

func (T *AddTask) GetTitle() string {
	return "广告配置"
}

