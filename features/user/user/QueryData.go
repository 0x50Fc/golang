package user

type QueryData struct {
	Items	[]*User	`json:"items,omitempty" name:"items" title:"用户"`
	Page	*QueryDataPage	`json:"page,omitempty" name:"page" title:"分页"`
}

