package top

type Page struct {
	P	int32	`json:"p" name:"p" title:"分页位置"`
	N	int32	`json:"n" name:"n" title:"单页记录数"`
	Count	int32	`json:"count" name:"count" title:"总页数"`
	Total	int32	`json:"total" name:"total" title:"总记录数"`
	TopId	int64	`json:"topId" name:"topid" title:"顶部ID"`
}

