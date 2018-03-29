package template

import (
	"github.com/go-wechat/util/httper"
)

const SendTemplate string = "https://api.weixin.qq.com/cgi-bin/message/template/send"

type TemplateMiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type SendTemplateReq struct {
	Touser      string                       `json:"touser"`
	TemplateID  string                       `json:"template_id"`
	Url         string                       `json:"url,omitempty"`
	MiniProgram *TemplateMiniProgram         `json:"miniprogram,omitempty"`
	Data        map[string]*TemplateDataItem `json:"data"`
}

type SendTemplateRes struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   int    `json:"msgid"`
}

// 发送模板消息
func SendTemplate(client httper.HttpClient, request *SendTemplateReq) *SendTemplateRes {
	var response = &SendTemplateRes{}
	return response
}
