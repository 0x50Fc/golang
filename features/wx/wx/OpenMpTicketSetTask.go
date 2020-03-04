package wx

type OpenMpTicketSetTask struct {
	Ticket	string	`json:"ticket" name:"ticket" title:"Ticket"`
}

func (T *OpenMpTicketSetTask) GetName() string {
	return "open/mp/ticket/set.json"
}

func (T *OpenMpTicketSetTask) GetTitle() string {
	return "开发平台 公众号授权 更新 Ticket"
}

