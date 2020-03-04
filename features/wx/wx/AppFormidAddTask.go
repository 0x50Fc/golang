package wx

type AppFormidAddTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Openid	string	`json:"openid" name:"openid" title:"openid"`
	Items	string	`json:"items" name:"items" title:"JSON\n[\n   {\"formid\":\"123\",\"etime\":1}\n]"`
}

func (T *AppFormidAddTask) GetName() string {
	return "app/formid/add.json"
}

func (T *AppFormidAddTask) GetTitle() string {
	return "小程序 添加 formid"
}

