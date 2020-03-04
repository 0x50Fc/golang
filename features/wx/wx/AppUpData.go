package wx

type AppUpData struct {
	Type	string	`json:"type" name:"type" title:"文件类型"`
	MediaId	string	`json:"media_id" name:"media_id" title:"媒体标示"`
	CreatedAt	int64	`json:"created_at" name:"created_at" title:"创建时间"`
}

