package wx

import (
	"github.com/hailongz/golang/db"
)

type FormId struct {
	db.Object
	Appid  string `json:"appid" name:"appid" title:"appid" length:"64" index:"ASC"`
	Openid string `json:"openid" name:"openid" title:"openid" length:"128" index:"ASC"`
	Formid string `json:"formid" name:"formid" title:"formid" length:"128"`
	Etime  int64  `json:"etime" name:"etime" title:"过期时间" index:"DESC"`
}

func (O *FormId) GetName() string {
	return "formid"
}

func (O *FormId) GetTitle() string {
	return "FormId"
}
