package article

type OutboxQueryData struct {
	Items	[]*Outbox	`json:"items,omitempty" name:"items" title:"发件"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

