package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopSkilsListAPIRequest
获取硬件平台设备下挂载的技能列表 API请求
taobao.ailab.aicloud.top.skils.list

提供给在硬件平台接入设备的技能列表 */
type TaobaoAilabAicloudTopSkilsListAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// query(模糊匹配skillName)
	_query string
	// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
	_type string
	// pageNo
	_pageNo int64
	// pageSize
	_pageSize int64
}

// NewTaobaoAilabAicloudTopSkilsListRequest 初始化TaobaoAilabAicloudTopSkilsListAPIRequest对象
func NewTaobaoAilabAicloudTopSkilsListRequest() *TaobaoAilabAicloudTopSkilsListAPIRequest {
	return &TaobaoAilabAicloudTopSkilsListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.skils.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetExt() string {
	return r._ext
}

// Set is Query Setter
// query(模糊匹配skillName)
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetQuery() string {
	return r._query
}

// Set is Type Setter
// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetType() string {
	return r._type
}

// Set is PageNo Setter
// pageNo
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// pageSize
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
