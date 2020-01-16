package app

import (
	"github.com/hailongz/golang/db"
)

type App struct {
	db.Object
	Title   string      `json:"title" name:"title" title:"标题" length:"255"`
	Uid     int64       `json:"uid" name:"uid" title:"用户ID"`
	Ver     int32       `json:"ver" name:"ver" title:"默认版本号"`
	LastVer int32       `json:"lastVer" name:"lastver" title:"最新版本号"`
	Options interface{} `json:"options,omitempty" name:"options" title:"其他数据" length:"-1"`
	Ctime   int64       `json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *App) GetName() string {
	return "app"
}

func (O *App) GetTitle() string {
	return "App"
}
