package wx

type AppFormidUseTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Openid	string	`json:"openid" name:"openid" title:"openid"`
}

func (T *AppFormidUseTask) GetName() string {
	return "app/formid/use.json"
}

func (T *AppFormidUseTask) GetTitle() string {
	return "小程序 使用 formid"
}

