package template

import (
	"github.com/go-wechat/util/httper"
)

const (
	SetIndustryUrl    string = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry"
	GetIndustryUrl    string = "https://api.weixin.qq.com/cgi-bin/template/get_industry"
	AddTemplateUrl    string = "https://api.weixin.qq.com/cgi-bin/template/api_add_template"
	GetTemplatesUrl   string = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template"
	DeleteTemplateUrl string = "https://api.weixin.qq.com/cgi-bin/template/del_private_template"
)

// 消息模板结构类型
type Template struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type SetIndustryReq struct {
	IndustryID1 string `json:"industry_id1"` // 公众号模板消息所属行业编号
	IndustryID2 string `json:"industry_id2"` // 公众号模板消息所属行业编号
}
type SetIndustryRes struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
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
func GetInsdustry(client httper.HttpClient, request *GetIndustryReq) *GetIndustryRes {

}

// 获得模板ID
type AddTemplateReq struct {
	TempalteIDShort string `json:"tempalte_id_short"`
}
type AddTemplateRes struct {
	Errcode    string `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	TemplateID string `json:"template_id"`
}

// 获取模板列表
type GetTemplatesRes struct {
	TemplateList []Template `json:"template_list"`
}

// 删除模板
type DeleteTemplateReq struct {
	Template_ID string `json:"template_id"`
}
type DeleteTemplateRes struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
