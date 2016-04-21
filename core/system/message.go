package system

import "encoding/json"

type RawMessage struct {
	BaseMessageType string `json:"type" form:"type"`
	Message         json.RawMessage
}

type Message struct {
	MessageType string `json:"type" form:"type"`
	Label       string `json:"label" form:"label"`
	Id          string `json:"id" form:"id"`
	Candidate   string `json:"candidate" form:"candidate"`
	Room        string `json:"room" form:"room"`
}

type Offer struct {
	MessageType string `json:"type" form:"type"`
	SDP         string `json:"sdp" form:"sdp"`
}
