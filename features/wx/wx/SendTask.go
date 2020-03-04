package wx

type SendTask struct {
	Uid	interface{}	`json:"uid,omitempty" title:"用户ID"`
	Uniqueid	interface{}	`json:"uniqueid,omitempty" title:"uniqueid"`
	Type	string	`json:"type,omitempty" title:"消息类型"`
	Body	string	`json:"body,omitempty" title:"消息内容:\nMessageType.Text:\n{\n   \"content\":\"Hello World\"\n}\nMessageType.Image:\nMessageType.Voice:\n{\n   \"media_id\":\"MEDIA_ID\"\n}\nMessageType.Video:\n{\n   \"media_id\":\"MEDIA_ID\",\n   \"thumb_media_id\":\"MEDIA_ID\",\n   \"title\":\"TITLE\",\n   \"description\":\"DESCRIPTION\"\n}\nMessageType.Music:\n{\n   \"title\":\"MUSIC_TITLE\",\n   \"description\":\"MUSIC_DESCRIPTION\",\n   \"musicurl\":\"MUSIC_URL\",\n   \"hqmusicurl\":\"HQ_MUSIC_URL\",\n   \"thumb_media_id\":\"THUMB_MEDIA_ID\" \n}\nMessageType.News:\n{\n   \"articles\": [\n   {\n       \"title\":\"Happy Day\",\n       \"description\":\"Is Really A Happy Day\",\n       \"url\":\"URL\",\n       \"picurl\":\"PIC_URL\"\n   }\n   ]\n}\nMessageType.Mpnews:\n{\n   \"media_id\":\"MEDIA_ID\"\n}\nMessageType.Msgmenu:\n{\n   \"head_content\": \"您对本次服务是否满意呢? \"\n  \"list\": [\n  {\n      \"id\": \"101\",\n      \"content\": \"满意\"\n  },\n  {\n      \"id\": \"102\",\n      \"content\": \"不满意\"\n  }\n  ],\n  \"tail_content\": \"欢迎再次光临\"\n}\nMessageType.Wxcard:\n{           \n  \"card_id\":\"123dsdajkasd231jhksad\"        \n}\nMessageType.Miniprogrampage:\n{\n      \"title\":\"title\",\n      \"appid\":\"appid\",\n      \"pagepath\":\"pagepath\",\n      \"thumb_media_id\":\"thumb_media_id\"\n}\nMessageType.Template:\n{\n   \"template_id\":\"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY\",\n     \"url\":\"http://weixin.qq.com/download\",  \n     \"miniprogram\":{\n       \"appid\":\"xiaochengxuappid12345\",\n       \"pagepath\":\"index?foo=bar\"\n     },          \n     \"data\":{\n             \"first\": {\n                 \"value\":\"恭喜你购买成功！\",\n                 \"color\":\"#173177\"\n             },\n             \"keyword1\":{\n                 \"value\":\"巧克力\",\n                 \"color\":\"#173177\"\n             },\n             \"keyword2\": {\n                 \"value\":\"39.8元\",\n                 \"color\":\"#173177\"\n             },\n             \"keyword3\": {\n                 \"value\":\"2014年9月22日\",\n                 \"color\":\"#173177\"\n             },\n             \"remark\":{\n                 \"value\":\"欢迎再次购买！\",\n                 \"color\":\"#173177\"\n             }\n     }\n}"`
	KfAccount	interface{}	`json:"kf_account,omitempty" title:"客服账号"`
}

func (T *SendTask) GetName() string {
	return "send.json"
}

func (T *SendTask) GetTitle() string {
	return "发送消息"
}

