package authority

import (
	"github.com/hailongz/golang/db"
)

type Role struct {
	db.Object
	Name	string	`json:"name" name:"name" title:"角色名" length:"64" unique:"ASC"`
	Title	string	`json:"title" name:"title" title:"角色说明" length:"255"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加" length:"-1"`
}

func (O *Role) GetName() string {
	return "role"
}

func (O *Role) GetTitle() string {
	return "角色"
}

