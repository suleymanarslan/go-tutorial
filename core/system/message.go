package system

type Message struct {
	MessageType     string `json:"type" form:"type"`
	Label string `json:"label" form:"label"`
	Id string `json:"id" form:"id"`
	Candidate string `json:"candidate" form:"candidate"`
	Room string `json:"room" form:"room"`

}
