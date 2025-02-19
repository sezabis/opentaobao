package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessPictureCheckAPIRequest 无线开放图片内容安全检查 API请求
// taobao.wireless.picture.check
//
// 无线开放内容检查，提供涉黄暴力政治图片检查。更详情介绍见 &lt;a href=&#34;https://help.aliyun.com/document_detail/70292.html&#34; target=&#34;blank&#34;&gt;阿里云内容安全&lt;/a&gt;
// 此API会进行两个场景审核，平均RT为1s。
type TaobaoWirelessPictureCheckAPIRequest struct {
	model.Params
	// 图片的URL,URL必须为淘系安全域名地址。图片格式支持png,jpg,webp
	_url string
}

// NewTaobaoWirelessPictureCheckRequest 初始化TaobaoWirelessPictureCheckAPIRequest对象
func NewTaobaoWirelessPictureCheckRequest() *TaobaoWirelessPictureCheckAPIRequest {
	return &TaobaoWirelessPictureCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessPictureCheckAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.picture.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessPictureCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUrl is Url Setter
// 图片的URL,URL必须为淘系安全域名地址。图片格式支持png,jpg,webp
func (r *TaobaoWirelessPictureCheckAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaoWirelessPictureCheckAPIRequest) GetUrl() string {
	return r._url
}
