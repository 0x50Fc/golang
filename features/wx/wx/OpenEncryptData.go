package wx

type OpenEncryptData struct {
	Nonce	string	`json:"nonce" name:"nonce" title:"nonce"`
	Timestamp	string	`json:"timestamp" name:"timestamp" title:"timestamp"`
	Signature	string	`json:"signature" name:"signature" title:"签名"`
	Content	string	`json:"content" name:"content" title:"编码后内容"`
}

