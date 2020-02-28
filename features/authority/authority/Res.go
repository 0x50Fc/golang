package authority

import (
	"github.com/hailongz/golang/db"
)

type Res struct {
	db.Object
	Path	string	`json:"path" name:"path" title:"资源路径" length:"128" unique:"ASC"`
	Title	string	`json:"title" name:"title" title:"说明" length:"255"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加" length:"-1"`
}

func (O *Res) GetName() string {
	return "res"
}

func (O *Res) GetTitle() string {
	return "资源"
}

