package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫根据规则增量更新商品 APIResponse
tmall.item.schema.increment.update

增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
*/
type TmallItemSchemaIncrementUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_schema_increment_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回商品发布结果
    
    UpdateItemResult   string `json:"update_item_result,omitempty" xml:"