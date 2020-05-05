package lookin

type AddData struct {
	Code	string	`json:"code" name:"code" title:""`
	Items	[]*Lookin	`json:"items,omitempty" name:"items" title:"" jsonType:"true"`
}

