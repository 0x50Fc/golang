package wx

type AppQrTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Scene	string	`json:"scene" name:"scene" title:"最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）"`
	Page	interface{}	`json:"page,omitempty" name:"page" title:"必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面"`
	Width	interface{}	`json:"width,omitempty" name:"width" title:"二维码的宽度，单位 px，最小 280px，最大 1280px"`
}

func (T *AppQrTask) GetName() string {
	return "app/qr.json"
}

func (T *AppQrTask) GetTitle() string {
	return "获取小程序二维码"
}

