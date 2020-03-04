package wx

type AppLoginTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Code	string	`json:"code" name:"code" title:"js_code"`
	EncryptedData	interface{}	`json:"encryptedData,omitempty" name:"encrypteddata" title:"用户信息编码数据"`
	Iv	interface{}	`json:"iv,omitempty" name:"iv" title:"加密算法的初始向量"`
}

func (T *AppLoginTask) GetName() string {
	return "app/login.json"
}

func (T *AppLoginTask) GetTitle() string {
	return "小程序登录"
}

