package wx

type MpRecvTask struct {
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

func (T *MpRecvTask) GetName() string {
	return "mp/recv.json"
}

func (T *MpRecvTask) GetTitle() string {
	return "公众号接收消息"
}

