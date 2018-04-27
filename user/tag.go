package user

import (
	"github.com/go-wechat/util/httper"
	"github.com/go-wechat/mp/common"
)

const (
	CreateTagsUrl string = "https://api.weixin.qq.com/cgi-bin/tags/create"
	GetTagsUrl string = "https://api.weixin.qq.com/cgi-bin/tags/get"
)

type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Count string `json:"count"`
}
// 创建用户标签
type TagName struct {
	Name string `json:"name"`
}
type CreateTagReq struct {
	Tag TagName `json:"tag"`
}
type CreateTagRes struct {
	common.CommonResponse
	Tag TagID `json:"tag"`
}
func CreateTag(client httper.HttpClient, request *CreateTagReq) (*CreateTagRes, error){
	var response = &CreateTagRes{}
	return response, nil
}

// 获取公众号已创建的标签
type GetTagsRes struct {
	Tags []Tag `json:"tags"`
}
func GetTags(client httper.HttpClient) (*GetTagsRes, error) {
	var response = &GetTagsRes{}
	return response, nil
}

// 编辑标签
type UpdateTagReq struct {

}
type UpdateTagRes struct {
	common.CommonResponse
}
func UpdateTag(client httper.HttpClient, request *UpdateTagReq) (*UpdateTagRes, error) {
	var response = &UpdateTagRes{}
	return response, nil
}

// 删除标签
type DeleteTagReq struct {

}
type DeleteTagRes struct {

}
func DeleteTag(client httper.HttpClient, request *DeleteTagReq) (*DeleteTagRes, error) {
	var response = &DeleteTagRes{}
	return response, nil
}
