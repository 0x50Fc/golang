package job

type AppQueryData struct {
	Items	[]*App	`json:"items,omitempty" name:"items" title:"应用"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

