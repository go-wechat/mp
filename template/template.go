package template

import (
	"github.com/go-wechat/util/httper"
	"github.com/go-wechat/official-account/common"
)

const (
	SetIndustryUrl    string = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"
	GetIndustryUrl    string = "https://api.weixin.qq.com/cgi-bin/template/get_industry"
	AddTemplateUrl    string = "https://api.weixin.qq.com/cgi-bin/template/api_add_template"
	GetTemplatesUrl   string = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	DeleteTemplateUrl string = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
	SendTemplateUrl   string = "https://api.weixin.qq.com/cgi-bin/message/template/send"
)

type SetIndustryReq struct {
	IndustryID1 string `json:"industry_id1"` // 公众号模板消息所属行业编号
	IndustryID2 string `json:"industry_id2"` // 公众号模板消息所属行业编号
}
type SetIndustryRes struct {
	common.CommonResponse
}
// 设置所属行业
func SetIndustry(client httper.HttpClient, request *SetIndustryReq) *SetIndustryRes{
	var response = &SendTemplateRes{}
	return response
}

// 获取设置的行业信息
type GetIndustryRes struct {
	PrimaryIndustry   string `json:"primary_industry"`
	SecondaryIndustry string `json:"secondary_industry"`
}
func GetInsdustry(client httper.HttpClient, request *GetIndustryReq) (*GetIndustryRes, error) {
	var response = &GetTemplatesRes{}
	return response, nil
}

// 获得模板ID
type AddTemplateReq struct {
	TempalteIDShort string `json:"tempalte_id_short"`
}
type AddTemplateRes struct {
	common.CommonResponse
	TemplateID string `json:"template_id"`
}
func AddTemplate(client httper.HttpClient, request *AddTemplateReq) (*AddTemplateRes, error) {
	var response = &AddTemplateRes{}
	return response, nil
}

// 获取模板列表
type Template struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}
type GetTemplatesRes struct {
	TemplateList []Template `json:"template_list"`
}
func GetTemplates(client httper.HttpClient) (*GetTemplatesRes, error){
	var response = &GetTemplatesRes{}
	return response, nil
}

// 删除模板
type DeleteTemplateReq struct {
	Template_ID string `json:"template_id"`
}
type DeleteTemplateRes struct {
	common.CommonResponse
}
func DeleteTemplate(client httper.HttpClient, request *DeleteTemplateReq) (*DeleteTemplateRes, error){
	var response = &DeleteTemplateRes{}
	return response, nil
}

// 发送模板消息
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
	common.CommonResponse
	Msgid   int    `json:"msgid"`
}
func SendTemplate(client httper.HttpClient, request *SendTemplateReq) (*SendTemplateRes, error) {
	var response = &SendTemplateRes{}
	return response, nil
}