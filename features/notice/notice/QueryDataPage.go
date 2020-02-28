package notice

type QueryDataPage struct {
	P	int32	`json:"p,omitempty" title:"分页位置"`
	N	int32	`json:"n,omitempty" title:"单页记录数"`
	Count	int32	`json:"count,omitempty" title:"总页数"`
	Total	int32	`json:"total,omitempty" title:"总记录数"`
}

