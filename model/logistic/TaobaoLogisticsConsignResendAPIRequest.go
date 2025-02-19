package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsConsignResendAPIRequest 修改物流公司和运单号 API请求
// taobao.logistics.consign.resend
//
// 支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：&lt;br&gt;
// 1、必须是已发货订单，自己联系发货的必须50天内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；
// 2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。
type TaobaoLogisticsConsignResendAPIRequest struct {
	model.Params
	// 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！
	_subTid []int64
	// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
	_outSid string
	// 物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取。<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font>
	_companyCode string
	// feature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
	_feature string
	// 商家的IP地址
	_sellerIp string
	// 淘宝交易ID
	_tid int64
	// 表明是否是拆单，默认值0，1表示拆单
	_isSplit int64
}

// NewTaobaoLogisticsConsignResendRequest 初始化TaobaoLogisticsConsignResendAPIRequest对象
func NewTaobaoLogisticsConsignResendRequest() *TaobaoLogisticsConsignResendAPIRequest {
	return &TaobaoLogisticsConsignResendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsConsignResendAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.consign.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsConsignResendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubTid is SubTid Setter
// 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！
func (r *TaobaoLogisticsConsignResendAPIRequest) SetSubTid(_subTid []int64) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetSubTid() []int64 {
	return r._subTid
}

// SetOutSid is OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsConsignResendAPIRequest) SetOutSid(_outSid string) error {
	r._outSid = _outSid
	r.Set("out_sid", _outSid)
	return nil
}

// GetOutSid OutSid Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetOutSid() string {
	return r._outSid
}

// SetCompanyCode is CompanyCode Setter
// 物流公司代码.如&#34;POST&#34;代表中国邮政,&#34;ZJS&#34;代表宅急送。调用 taobao.logistics.companies.get 获取。&lt;br&gt;&lt;font color=&#39;red&#39;&gt;如果是货到付款订单，选择的物流公司必须支持货到付款发货方式&lt;/font&gt;
func (r *TaobaoLogisticsConsignResendAPIRequest) SetCompanyCode(_companyCode string) error {
	r._companyCode = _companyCode
	r.Set("company_code", _companyCode)
	return nil
}

// GetCompanyCode CompanyCode Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetCompanyCode() string {
	return r._companyCode
}

// SetFeature is Feature Setter
// feature参数格式&lt;br&gt;范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B&lt;br&gt;identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔&lt;br&gt;“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。&#34;|&#34;不同商品间的分隔符。&lt;br&gt;例1商品和2商品，之间就用&#34;|&#34;分开。&lt;br&gt;TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)&lt;br&gt;&#34;:&#34;TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。&lt;br&gt;&#34;,&#34; 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。&lt;br&gt;具体:当订单中A商品的数量为2个，其中手机串号分别为&#34;12345&#34;,&#34;67890&#34;。&lt;br&gt;参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的&#34;12345&#34;,&#34;67890&#34;.说明本订单A宝贝的数量为2，值分别为&#34;12345&#34;,&#34;67890&#34;。&lt;br&gt;当存在&#34;|&#34;时，就说明订单中存在多个商品，商品间用&#34;|&#34;分隔了开来。|&#34;之后的内容含义同上。
func (r *TaobaoLogisticsConsignResendAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetFeature() string {
	return r._feature
}

// SetSellerIp is SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsConsignResendAPIRequest) SetSellerIp(_sellerIp string) error {
	r._sellerIp = _sellerIp
	r.Set("seller_ip", _sellerIp)
	return nil
}

// GetSellerIp SellerIp Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetSellerIp() string {
	return r._sellerIp
}

// SetTid is Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsConsignResendAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetTid() int64 {
	return r._tid
}

// SetIsSplit is IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsConsignResendAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// GetIsSplit IsSplit Getter
func (r TaobaoLogisticsConsignResendAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}
