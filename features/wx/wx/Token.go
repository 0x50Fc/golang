package wx

import (
	"github.com/hailongz/golang/db"
)

type Token struct {
	db.Object
	Type	int32	`json:"type" name:"type" title:"类型" index:"ASC"`
	Appid	string	`json:"appid" name:"appid" title:"appid" length:"64" index:"ASC"`
	AccessToken	string	`json:"-" name:"access_token" title:"access_token" length:"255"`
	Etime	int64	`json:"etime" name:"etime" title:"过期时间"`
}

func (O *Token) GetName() string {
	return "token"
}

func (O *Token) GetTitle() string {
	return "Token"
}

