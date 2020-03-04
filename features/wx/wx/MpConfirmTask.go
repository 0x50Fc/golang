package wx

type MpConfirmTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"state"`
	Code	string	`json:"code" name:"code" title:"code"`
}

func (T *MpConfirmTask) GetName() string {
	return "mp/confirm.json"
}

func (T *MpConfirmTask) GetTitle() string {
	return "授权确认"
}

