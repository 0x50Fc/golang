package inbox

type InboxQueryData struct {
	Items	[]*Inbox	`json:"items,omitempty" name:"items" title:"收件"`
	Page	*TopPage	`json:"page,omitempty" name:"page" title:"分页"`
}

