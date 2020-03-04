package wx

import (
	"github.com/hailongz/golang/db"
)

type User struct {
	db.Object
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Type	int32	`json:"type" name:"type" title:"类型" index:"ASC"`
	Appid	string	`json:"appid" name:"appid" title:"appid" length:"64" index:"ASC"`
	Openid	string	`json:"openid" name:"openid" title:"openid" length:"128" index:"ASC"`
	Unionid	string	`json:"unionid" name:"unionid" title:"unionid" length:"128" index:"ASC"`
	AccessToken	string	`json:"access_token" name:"access_token" title:"access_token" length:"255"`
	RefreshToken	string	`json:"refresh_token" name:"refresh_token" title:"refresh_token" length:"255"`
	SessionKey	string	`json:"session_key" name:"session_key" title:"session_key" length:"128"`
	Nick	string	`json:"nick" name:"nick" title:"昵称" length:"255"`
	Logo	string	`json:"logo" name:"logo" title:"头像" length:"2048"`
	Country	string	`json:"country" name:"country" title:"国家" length:"64"`
	Lang	string	`json:"lang" name:"lang" title:"语言" length:"64"`
	Province	string	`json:"province" name:"province" title:"省份" length:"64"`
	City	string	`json:"city" name:"city" title:"城市" length:"64"`
	Gender	int32	`json:"gender" name:"gender" title:"性别"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
	Etime	int64	`json:"etime" name:"etime" title:"过期时间"`
	State	int32	`json:"state" name:"state" title:"关注状态"`
	Mtime	int64	`json:"mtime" name:"mtime" title:"最后绑定时间"`
}

func (O *User) GetName() string {
	return "user"
}

func (O *User) GetTitle() string {
	return "用户"
}

