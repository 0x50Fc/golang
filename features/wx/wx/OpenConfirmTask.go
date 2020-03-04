package wx

type OpenConfirmTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"state"`
	Code	string	`json:"code" name:"code" title:"code"`
}

func (T *OpenConfirmTask) GetName() string {
	return "open/confirm.json"
}

func (T *OpenConfirmTask) GetTitle() string {
	return "开发平台授权确认"
}

