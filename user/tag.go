package user

import (
	"github.com/go-wechat/util/httper"
)

const (
	CreateTagsUrl string = "https://api.weixin.qq.com/cgi-bin/tags/create"
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
	Tag TagID `json:"tag"`
}
func CreateTag(client httper.HttpClient, request *CreateTagReq) *CreateTagRes{
	var response = &CreateTagRes{}
	return response
}
