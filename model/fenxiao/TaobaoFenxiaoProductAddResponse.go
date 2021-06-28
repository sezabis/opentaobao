package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加产品 APIResponse
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
*/
type TaobaoFenxiaoProductAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品ID
    
    Pid   int64 `json:"pid,omitempty" xml:"