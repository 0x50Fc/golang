package notice

type QueryData struct {
	Items	[]*Notice	`json:"items,omitempty" name:"items" title:"通知"`
	Page	*TopPage	`json:"page,omitempty" name:"page" title:"分页"`
}

