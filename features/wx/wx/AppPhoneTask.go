package wx

type AppPhoneTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Openid	interface{}	`json:"openid,omitempty" name:"openid" title:"openid"`
	SessionKey	interface{}	`json:"session_key,omitempty" name:"session_key" title:"session_key"`
	EncryptedData	string	`json:"encryptedData" name:"encrypteddata" title:"用户信息编码数据"`
	Iv	string	`json:"iv" name:"iv" title:"加密算法的初始向量"`
}

func (T *AppPhoneTask) GetName() string {
	return "app/phone.json"
}

func (T *AppPhoneTask) GetTitle() string {
	return "小程序获取手机号"
}

