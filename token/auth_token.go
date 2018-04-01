package token

/**
网页授权access_token
微信网页授权是通过OAuth2.0机制实现的, 在用户授权给公众号后, 公众号可以获取到一个网页授权特有的接口调用凭证(网页授权access_token).
通过网页授权access_token可以进行授权后接口调用，如获取用户基本信息.
*/


type ApiToken string

func (token *ApiToken) GetToken() string {

}

func (token *ApiToken)  RefreshToken() string {

}
