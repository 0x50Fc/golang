package doc

type DocQueryData struct {
	Items	[]*Doc	`json:"items,omitempty" name:"items" title:"文档"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

