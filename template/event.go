package template

const (
	TemplateEventType string = "TEMPLATESENDJOBFINISH"
	TemplateEventStatusList = []string{ "success", "failed:user block", "failed: system failed"}
)

type TemplateEvent struct {
	ToUserName string `xml:"ToUserName",json:"ToUserName"`
	FromUserName string `xml:"FromUserName",json:"FromUserName"`
	CreateTime int `xml:"CreateTime",json:"CreateTime"`
	MsgType string `xml:"MsgType",json:"MsgType"`
	Event string `xml:"Event",json:"Event"`
	MsgID int `xml:"MsgId",json:"MsgID"`
	Status string `xml:"Status",json:"Status"`
}
