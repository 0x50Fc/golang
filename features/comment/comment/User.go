package comment

type User struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID" index:"ASC"`
	Count	int32	`json:"count" name:"count" title:"发布的数量"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"最后发布时间"`
}

