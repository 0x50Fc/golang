package wx

type AppSendTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Openid	string	`json:"openid" name:"openid" title:"openid"`
	Type	string	`json:"type" name:"type" title:"消息类型"`
	Body	string	`json:"body" name:"body" title:"消息内容:\nMessageType.Text:\n{\n   \"content\":\"Hello World\"\n}\nMessageType.Image:\nMessageType.Voice:\n{\n   \"media_id\":\"MEDIA_ID\"\n}\nMessageType.Video:\n{\n   \"media_id\":\"MEDIA_ID\",\n   \"thumb_media_id\":\"MEDIA_ID\",\n   \"title\":\"TITLE\",\n   \"description\":\"DESCRIPTION\"\n}\nMessageType.Music:\n{\n   \"title\":\"MUSIC_TITLE\",\n   \"description\":\"MUSIC_DESCRIPTION\",\n   \"musicurl\":\"MUSIC_URL\",\n   \"hqmusicurl\":\"HQ_MUSIC_URL\",\n   \"thumb_media_id\":\"THUMB_MEDIA_ID\" \n}\nMessageType.News:\n{\n   \"articles\": [\n   {\n       \"title\":\"Happy Day\",\n       \"description\":\"Is Really A Happy Day\",\n       \"url\":\"URL\",\n       \"picurl\":\"PIC_URL\"\n   }\n   ]\n}\nMessageType.Mpnews:\n{\n   \"media_id\":\"MEDIA_ID\"\n}\nMessageType.Msgmenu:\n{\n   \"head_content\": \"您对本次服务是否满意呢? \"\n  \"list\": [\n  {\n      \"id\": \"101\",\n      \"content\": \"满意\"\n  },\n  {\n      \"id\": \"102\",\n      \"content\": \"不满意\"\n  }\n  ],\n  \"tail_content\": \"欢迎再次光临\"\n}\nMessageType.Wxcard:\n{           \n  \"card_id\":\"123dsdajkasd231jhksad\"        \n}\nMessageType.Miniprogrampage:\n{\n      \"title\":\"title\",\n      \"appid\":\"appid\",\n      \"pagepath\":\"pagepath\",\n      \"thumb_media_id\":\"thumb_media_id\"\n}\nMessageType.Template:\n{\n  \"touser\": \"OPENID\",\n  \"template_id\": \"TEMPLATE_ID\",\n  \"page\": \"index\",\n  \"form_id\": \"FORMID\",\n  \"data\": {\n      \"keyword1\": {\n          \"value\": \"339208499\"\n      },\n      \"keyword2\": {\n          \"value\": \"2015年01月05日 12:30\"\n      },\n      \"keyword3\": {\n          \"value\": \"腾讯微信总部\"\n      } ,\n      \"keyword4\": {\n          \"value\": \"广州市海珠区新港中路397号\"\n      }\n  },\n  \"emphasis_keyword\": \"keyword1.DATA\"\n}"`
	KfAccount	interface{}	`json:"kf_account,omitempty" name:"kf_account" title:"客服账号"`
}

func (T *AppSendTask) GetName() string {
	return "app/send.json"
}

func (T *AppSendTask) GetTitle() string {
	return "发送消息"
}

