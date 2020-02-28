package settings

import (
	"github.com/hailongz/golang/db"
)

type Setting struct {
	db.Object
	Name	string	`json:"name" name:"name" title:"设置名" length:"128" index:"ASC"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加" length:"-1"`
}

func (O *Setting) GetName() string {
	return "setting"
}

func (O *Setting) GetTitle() string {
	return "系统设置"
}

