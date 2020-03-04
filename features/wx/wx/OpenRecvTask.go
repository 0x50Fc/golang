package wx

type OpenRecvTask struct {
	Appid	string	`json:"appid" title:"appid"`
	Type	interface{}	`json:"type,omitempty" title:"内容类型 json/xml 默认 json"`
	EncodingKey	string	`json:"encodingKey" title:"encodingKey"`
	Echostr	string	`json:"echostr" title:"echostr"`
	Nonce	string	`json:"nonce" title:"nonce"`
	Timestamp	string	`json:"timestamp" title:"timestamp"`
	Signature	string	`json:"signature" title:"signature"`
	Token	string	`json:"token" title:"token"`
	Content	string	`json:"content" title:"内容"`
}

func (T *OpenRecvTask) GetName() string {
	return "open/recv.json"
}

func (T *OpenRecvTask) GetTitle() string {
	return "开放平台接收消息"
}

