package user

// OAuthOtherInfo 结构体
type OAuthOtherInfo struct {
	// access_token
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// avatarUrl
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 三方平台的openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 三方平台的unionId
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 三方平台类型
	PlatformType int64 `json:"platform_type,omitempty" xml:"platform_type,omitempty"`
}
