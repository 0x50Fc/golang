package wx

type ContentList struct {
	Items []*Content `json:"items,omitempty" name:"items" title:"内容"`
	Page  *Page      `json:"page,omitempty" name:"page" title:"分页"`
}
