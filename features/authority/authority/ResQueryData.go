package authority

type ResQueryData struct {
	Items	[]*Res	`json:"items,omitempty" name:"items" title:"资源"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

