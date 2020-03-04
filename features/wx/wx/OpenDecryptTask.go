package wx

type OpenDecryptTask struct {
	Token	string	`json:"token" name:"token" title:"Token"`
	EncodingKey	string	`json:"encodingKey" name:"encodingkey" title:"encodingKey"`
	Nonce	string	`json:"nonce" name:"nonce" title:"nonce"`
	Timestamp	string	`json:"timestamp" name:"timestamp" title:"timestamp"`
	Signature	string	`json:"signature" name:"signature" title:"签名"`
	Content	string	`json:"content" name:"content" title:"内容 XML"`
}

func (T *OpenDecryptTask) GetName() string {
	return "open/decrypt.json"
}

func (T *OpenDecryptTask) GetTitle() string {
	return "开放平台解码消息"
}

