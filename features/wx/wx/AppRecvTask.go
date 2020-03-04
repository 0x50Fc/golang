package wx

type AppRecvTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"内容类型 json/xml 默认 json"`
	EncodingKey	string	`json:"encodingKey" name:"encodingkey" title:"encodingKey"`
	Echostr	string	`json:"echostr" name:"echostr" title:"echostr"`
	Nonce	string	`json:"nonce" name:"nonce" title:"nonce"`
	Timestamp	string	`json:"timestamp" name:"timestamp" title:"timestamp"`
	Signature	string	`json:"signature" name:"signature" title:"signature"`
	Token	string	`json:"token" name:"token" title:"token"`
	Content	string	`json:"content" name:"content" title:"内容"`
}

func (T *AppRecvTask) GetName() string {
	return "app/recv.json"
}

func (T *AppRecvTask) GetTitle() string {
	return "接收消息"
}

