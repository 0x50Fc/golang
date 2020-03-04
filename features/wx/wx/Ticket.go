package wx

import (
	"github.com/hailongz/golang/db"
)

type Ticket struct {
	db.Object
	Type	string	`json:"type" name:"type" title:"类型" length:"32" index:"ASC"`
	Appid	string	`json:"appid" name:"appid" title:"appid" length:"64" index:"ASC"`
	Ticket	string	`json:"-" name:"ticket" title:"ticket" length:"255"`
	Etime	int64	`json:"etime" name:"etime" title:"过期时间"`
}

func (O *Ticket) GetName() string {
	return "ticket"
}

func (O *Ticket) GetTitle() string {
	return "Ticket"
}

