package auth

import (
	"github.com/hailongz/golang/db"
)

type Auth struct {
	db.Object
	Key	string	`json:"key,omitempty" title:"唯一键" length:"255" unique:"ASC"`
	Value	string	`json:"value,omitempty" title:"值" length:"-1"`
	Etime	int64	`json:"etime,omitempty" title:"失效时间"`
}

func (O *Auth) GetName() string {
	return "auth"
}

func (O *Auth) GetTitle() string {
	return "验证"
}

