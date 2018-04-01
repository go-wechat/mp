package user

import (
	"github.com/go-wechat/util/httper"
	"github.com/go-wechat/official-account/common"
)

const (
	CreateTagsUrl string = "https://api.weixin.qq.com/cgi-bin/tags/create"
	GetTagsUrl string = "https://api.weixin.qq.com/cgi-bin/tags/get"
)

// 创建用户标签
type TagName struct {
	Name string `json:"name"`
}
type CreateTagReq struct {
	Tag TagName `json:"tag"`
}
type TagID struct {
	ID string `json:"id"`
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
type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Count string `json:"count"`
}
type GetTagsRes struct {
	Tags []Tag `json:"tags"`
}
func GetTags(client httper.HttpClient) (*GetTagsRes, error) {
	var response = &GetTagsRes{}
	return response, nil
}