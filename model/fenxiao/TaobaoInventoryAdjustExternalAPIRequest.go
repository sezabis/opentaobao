package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryAdjustExternalAPIRequest 非交易库存调整单 API请求
// taobao.inventory.adjust.external
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalAPIRequest struct {
	model.Params
	// 商家外部定单号
	_bizUniqueCode string
	// 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
	_occupyOperateCode string
	// test
	_operateType string
	// 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
	_bizType string
	// 业务操作时间
	_operateTime string
	// 商家仓库编码
	_storeCode string
	// 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}]
	_items string
	// test
	_reduceType string
}

// NewTaobaoInventoryAdjustExternalRequest 初始化TaobaoInventoryAdjustExternalAPIRequest对象
func NewTaobaoInventoryAdjustExternalRequest() *TaobaoInventoryAdjustExternalAPIRequest {
	return &TaobaoInventoryAdjustExternalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryAdjustExternalAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.adjust.external"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryAdjustExternalAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizUniqueCode is BizUniqueCode Setter
// 商家外部定单号
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetBizUniqueCode(_bizUniqueCode string) error {
	r._bizUniqueCode = _bizUniqueCode
	r.Set("biz_unique_code", _bizUniqueCode)
	return nil
}

// GetBizUniqueCode BizUniqueCode Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetBizUniqueCode() string {
	return r._bizUniqueCode
}

// SetOccupyOperateCode is OccupyOperateCode Setter
// 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetOccupyOperateCode(_occupyOperateCode string) error {
	r._occupyOperateCode = _occupyOperateCode
	r.Set("occupy_operate_code", _occupyOperateCode)
	return nil
}

// GetOccupyOperateCode OccupyOperateCode Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetOccupyOperateCode() string {
	return r._occupyOperateCode
}

// SetOperateType is OperateType Setter
// test
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetBizType is BizType Setter
// 外部订单类型, BALANCE:盘点、NON_TAOBAO_TRADE:非淘宝交易、ALLOCATE:调拨、OTHERS:其他
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetBizType() string {
	return r._bizType
}

// SetOperateTime is OperateTime Setter
// 业务操作时间
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetOperateTime(_operateTime string) error {
	r._operateTime = _operateTime
	r.Set("operate_time", _operateTime)
	return nil
}

// GetOperateTime OperateTime Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetOperateTime() string {
	return r._operateTime
}

// SetStoreCode is StoreCode Setter
// 商家仓库编码
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetItems is Items Setter
// 商品初始库存信息： [{&#34;scItemId&#34;:&#34;商品后端ID，如果有传scItemCode,参数可以为0&#34;,&#34;scItemCode&#34;:&#34;商品商家编码&#34;,&#34;inventoryType&#34;:&#34;库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,&#34;quantity&#34;:&#34;数量(正数)&#34;}]
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetItems(_items string) error {
	r._items = _items
	r.Set("items", _items)
	return nil
}

// GetItems Items Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetItems() string {
	return r._items
}

// SetReduceType is ReduceType Setter
// test
func (r *TaobaoInventoryAdjustExternalAPIRequest) SetReduceType(_reduceType string) error {
	r._reduceType = _reduceType
	r.Set("reduce_type", _reduceType)
	return nil
}

// GetReduceType ReduceType Getter
func (r TaobaoInventoryAdjustExternalAPIRequest) GetReduceType() string {
	return r._reduceType
}
