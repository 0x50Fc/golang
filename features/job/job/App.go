package job

import (
	"github.com/hailongz/golang/db"
)

type App struct {
	db.Object
	Alias	string	`json:"alias" name:"alias" title:"别名" length:"128"`
	Type	string	`json:"type" name:"type" title:"类型" length:"128"`
	Content	string	`json:"content" name:"content" title:"内容" length:"-1"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *App) GetName() string {
	return "app"
}

func (O *App) GetTitle() string {
	return "应用"
}

