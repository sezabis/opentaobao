package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreReallocateAPIRequest rellocate API请求
// taobao.omniorder.store.reallocate
//
// 门店发货提供改派接口
type TaobaoOmniorderStoreReallocateAPIRequest struct {
	model.Params
	// 子订单号
	_subOrderIds []string
	// 电商仓code
	_warehouseCode string
	// 主订单号
	_mainOrderId int64
	// 门店Id
	_storeId int64
}

// NewTaobaoOmniorderStoreReallocateRequest 初始化TaobaoOmniorderStoreReallocateAPIRequest对象
func NewTaobaoOmniorderStoreReallocateRequest() *TaobaoOmniorderStoreReallocateAPIRequest {
	return &TaobaoOmniorderStoreReallocateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.reallocate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubOrderIds is SubOrderIds Setter
// 子订单号
func (r *TaobaoOmniorderStoreReallocateAPIRequest) SetSubOrderIds(_subOrderIds []string) error {
	r._subOrderIds = _subOrderIds
	r.Set("sub_order_ids", _subOrderIds)
	return nil
}

// GetSubOrderIds SubOrderIds Getter
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetSubOrderIds() []string {
	return r._subOrderIds
}

// SetWarehouseCode is WarehouseCode Setter
// 电商仓code
func (r *TaobaoOmniorderStoreReallocateAPIRequest) SetWarehouseCode(_warehouseCode string) error {
	r._warehouseCode = _warehouseCode
	r.Set("warehouse_code", _warehouseCode)
	return nil
}

// GetWarehouseCode WarehouseCode Getter
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetWarehouseCode() string {
	return r._warehouseCode
}

// SetMainOrderId is MainOrderId Setter
// 主订单号
func (r *TaobaoOmniorderStoreReallocateAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetStoreId is StoreId Setter
// 门店Id
func (r *TaobaoOmniorderStoreReallocateAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoOmniorderStoreReallocateAPIRequest) GetStoreId() int64 {
	return r._storeId
}
