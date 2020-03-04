package wx

type MpPullTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
}

func (T *MpPullTask) GetName() string {
	return "mp/pull.json"
}

func (T *MpPullTask) GetTitle() string {
	return "拉去关注微信公众号的用户"
}

