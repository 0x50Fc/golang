package authority

import (
	"github.com/hailongz/golang/db"
)

type Authority struct {
	db.Object
	Uid     int64       `json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	RoleId  int64       `json:"roleId" name:"roleid" title:"角色ID" index:"ASC"`
	ResId   int64       `json:"resId" name:"resid" title:"资源ID" index:"ASC"`
	Options interface{} `json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加" length:"-1"`
}

func (O *Authority) GetName() string {
	return "authority"
}

func (O *Authority) GetTitle() string {
	return "权限"
}
