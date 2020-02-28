package authority

type RoleQueryData struct {
	Items	[]*Role	`json:"items,omitempty" name:"items" title:"资源"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

