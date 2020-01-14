package inbox

type User struct {
	Uid   int64 `json:"uid" name:"uid" title:"用户ID"`
	Mid   int64 `json:"mid" name:"mid" title:"内容ID"`
	Iid   int64 `json:"iid" name:"iid" title:"内容项ID"`
	Ctime int64 `json:"ctime" name:"ctime" title:"最后时间"`
	Fuid  int64 `json:"fuid" name:"fuid" title:"发布者ID"`
}
