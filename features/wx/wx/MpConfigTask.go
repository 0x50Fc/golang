package wx

type MpConfigTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Noncestr	interface{}	`json:"noncestr,omitempty" name:"noncestr" title:"noncestr 不存在是自动生成"`
	Timestamp	interface{}	`json:"timestamp,omitempty" name:"timestamp" title:"noncestr 不存时是自动生成"`
	Url	string	`json:"url" name:"url" title:"签名URL"`
}

func (T *MpConfigTask) GetName() string {
	return "mp/config.json"
}

func (T *MpConfigTask) GetTitle() string {
	return "获取JS签名配置信息"
}

