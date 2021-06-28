package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家会员数据上传 APIResponse
alibaba.tcls.aelophy.merchant.user.upload

商家会员数据上传
*/
type AlibabaTclsAelophyMerchantUserUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_merchant_user_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    ApiResult   *AlibabaTclsAelophyMerchantUserUploadApiResult `json:"api_result,omitempty" xml:"