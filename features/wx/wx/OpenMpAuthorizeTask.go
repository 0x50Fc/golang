package wx

type OpenMpAuthorizeTask struct {
	OpenType	interface{}	`json:"openType,omitempty" name:"opentype" title:"授权方式"`
	AuthType	interface{}	`json:"authType,omitempty" name:"authtype" title:"授权类型"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"公众号/小程序 Appid"`
	RedirectUri	string	`json:"redirect_uri" name:"redirect_uri" title:"redirect_uri"`
}

func (T *OpenMpAuthorizeTask) GetName() string {
	return "open/mp/authorize.json"
}

func (T *OpenMpAuthorizeTask) GetTitle() string {
	return "开发平台 公众号授权 获取授权URL"
}

