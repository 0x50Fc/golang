package wx

import (
	"github.com/hailongz/golang/db"
)

type Open struct {
	db.Object
	Appid	string	`json:"appid" name:"appid" title:"appid" length:"64" index:"ASC"`
	Ticket	string	`json:"-" name:"ticket" title:"ticket" length:"255"`
	AccessToken	string	`json:"-" name:"access_token" title:"access_token" length:"255"`
	RefreshToken	string	`json:"-" name:"refresh_token" title:"refresh_token" length:"255"`
	Etime	int64	`json:"etime" name:"etime" title:"过期时间"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
}

func (O *Open) GetName() string {
	return "open"
}

func (O *Open) GetTitle() string {
	return "开发平台"
}

