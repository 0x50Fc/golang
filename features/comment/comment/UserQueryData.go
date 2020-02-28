package comment

type UserQueryData struct {
	Items	[]*User	`json:"items,omitempty" name:"items" title:""`
	Page	*Page	`json:"page,omitempty" name:"page" title:""`
}

