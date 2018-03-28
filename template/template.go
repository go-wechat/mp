package template

// 模板结构类型
type Template struct {
	TemplateID string `json:"template_id"`
	Title string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry string `json:"deputy_industry"`
	Content string `json:"content"`
	Example string `json:"example"`
}

// 设置所属行业
type SetIndustryReq struct {
	IndustryID1 string `json:"industry_id1"`
	IndustryID2 string `json:"industry_id2"`
}
type SetIndustryRes struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

// 获取设置的行业信息
type GetIndustryRes struct {
	PrimaryIndustry string `json:"primary_industry"`
	SecondaryIndustry string `json:"secondary_industry"`
}

// 获得模板ID
type AddTemplateReq struct {
	TempalteIDShort string `json:"tempalte_id_short"`
}
type AddTemplateRes struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
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
	Errmsg string `json:"errmsg"`
}

