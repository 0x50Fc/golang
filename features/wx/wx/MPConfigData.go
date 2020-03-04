package wx

type MPConfigData struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Timestamp	int64	`json:"timestamp" name:"timestamp" title:"timestamp"`
	NonceStr	string	`json:"nonceStr" name:"noncestr" title:"nonceStr"`
	Signature	string	`json:"signature" name:"signature" title:"signature"`
}

