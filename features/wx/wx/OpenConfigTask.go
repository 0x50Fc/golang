package wx

type OpenConfigTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Noncestr	interface{}	`json:"noncestr,omitempty" name:"noncestr" title:"noncestr 不存在是自动生成"`
	Timestamp	interface{}	`json:"timestamp,omitempty" name:"timestamp" title:"noncestr 不存时是自动生成"`
	Url	string	`json:"url" name:"url" title:"签名URL"`
}

func (T *OpenConfigTask) GetName() string {
	return "open/config.json"
}

func (T *OpenConfigTask) GetTitle() string {
	return "开发平台获取JS签名配置信息"
}

