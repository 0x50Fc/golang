package inbox

type UserQueryData struct {
	Items []*User `json:"items,omitempty" name:"items" title:"用户"`
	Page  *Page   `json:"page,omitempty" name:"page" title:"分页"`
}
