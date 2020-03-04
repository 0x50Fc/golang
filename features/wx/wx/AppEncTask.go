package wx

type AppEncTask struct {
	Appid	string	`json:"appid,omitempty" title:"appid"`
	Type	interface{}	`json:"type,omitempty" title:"内容类型 json/xml 默认 json"`
	EncodingKey	string	`json:"encodingKey,omitempty" title:"encodingKey"`
	Content	string	`json:"content,omitempty" title:"内容"`
}

func (T *AppEncTask) GetName() string {
	return "app/enc.json"
}

func (T *AppEncTask) GetTitle() string {
	return "编码回复消息"
}

