package wx

type MpAuthorizeTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"state"`
	Scope	interface{}	`json:"scope,omitempty" name:"scope" title:"scope"`
	RedirectUri	string	`json:"redirect_uri" name:"redirect_uri" title:"redirect_uri"`
}

func (T *MpAuthorizeTask) GetName() string {
	return "mp/authorize.json"
}

func (T *MpAuthorizeTask) GetTitle() string {
	return "获取授权URL"
}

