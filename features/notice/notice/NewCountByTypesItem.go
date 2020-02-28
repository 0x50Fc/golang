package notice

type NewCountByTypesItem struct {
	Type	int32	`json:"type" name:"type" title:"类型"`
	Count	int32	`json:"count" name:"count" title:"记录数"`
}

