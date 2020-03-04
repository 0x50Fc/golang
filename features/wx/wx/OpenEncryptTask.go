package wx

type OpenEncryptTask struct {
	Token	string	`json:"token" name:"token" title:"Token"`
	EncodingKey	string	`json:"encodingKey" name:"encodingkey" title:"encodingKey"`
	Nonce	interface{}	`json:"nonce,omitempty" name:"nonce" title:"nonce"`
	Timestamp	interface{}	`json:"timestamp,omitempty" name:"timestamp" title:"timestamp"`
	Content	string	`json:"content" name:"content" title:"内容 JSON"`
}

func (T *OpenEncryptTask) GetName() string {
	return "open/encrypt.json"
}

func (T *OpenEncryptTask) GetTitle() string {
	return "开放平台编码消息"
}

