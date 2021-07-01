package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMemberBindmemberAPIRequest
绑定会员 API请求
alibaba.mj.member.bindmember

用于绑定喵街数字化会员 */
type AlibabaMjMemberBindmemberAPIRequest struct {
	model.Params
	// 用户号
	_userId int64
	// 商城Id
	_mallId int64
	// open_id
	_openId string
	// 渠道
	_channel string
}

// NewAlibabaMjMemberBindmemberRequest 初始化AlibabaMjMemberBindmemberAPIRequest对象
func NewAlibabaMjMemberBindmemberRequest() *AlibabaMjMemberBindmemberAPIRequest {
	return &AlibabaMjMemberBindmemberAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberBindmemberAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.member.bindmember"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberBindmemberAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户号
func (r *AlibabaMjMemberBindmemberAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is MallId Setter
// 商城Id
func (r *AlibabaMjMemberBindmemberAPIRequest) SetMallId(_mallId int64) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// Get MallId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetMallId() int64 {
	return r._mallId
}

// Set is OpenId Setter
// open_id
func (r *AlibabaMjMemberBindmemberAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// Get OpenId Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetOpenId() string {
	return r._openId
}

// Set is Channel Setter
// 渠道
func (r *AlibabaMjMemberBindmemberAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r AlibabaMjMemberBindmemberAPIRequest) GetChannel() string {
	return r._channel
}
