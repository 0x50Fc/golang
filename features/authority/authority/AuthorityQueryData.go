package authority

type AuthorityQueryData struct {
	Items	[]*Authority	`json:"items,omitempty" name:"items" title:"授权"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

