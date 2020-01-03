package vcode

import (
	"github.com/hailongz/golang/db"
)

type VCode struct {
	db.Object
	Key	string	`json:"key" name:"key" title:"Key" length:"128" unique:"ASC"`
	Code	string	`json:"code" name:"code" title:"数字验证码 最大 12位" length:"12"`
	Hash	string	`json:"hash" name:"hash" title:"32位 HASH" length:"32"`
	Etime	int64	`json:"etime" name:"etime" title:"过期时间"`
}

func (O *VCode) GetName() string {
	return "vcode"
}

func (O *VCode) GetTitle() string {
	return "验证码"
}

