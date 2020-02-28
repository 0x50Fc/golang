package adv

type QueryData struct {
	Items	[]*Adv	`json:"items,omitempty" name:"items" title:"广告"`
	Page	*Page	`json:"page,omitempty" name:"page" title:"分页"`
}

