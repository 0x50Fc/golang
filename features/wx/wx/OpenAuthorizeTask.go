package wx

type OpenAuthorizeTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"state"`
	Scope	interface{}	`json:"scope,omitempty" name:"scope" title:"scope"`
	RedirectUri	string	`json:"redirect_uri" name:"redirect_uri" title:"redirect_uri"`
}

func (T *OpenAuthorizeTask) GetName() string {
	return "open/authorize.json"
}

func (T *OpenAuthorizeTask) GetTitle() string {
	return "开发平台获取授权URL"
}

