package jst

// CheckSignSkuResponse 结构体
type CheckSignSkuResponse struct {
	// NO_AUTH-您无权查询此商家信息;  HAS_SIGNED-已与当前服务商签约自助修改商品信息服务; NO_SIGNED-当前商家自助修改商品信息未签约服务商 HAS_SIGNED_OTHERS-已与其他服务商签约自助修改商品信息
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商家绑定的appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// app名称
	Apptitle string `json:"apptitle,omitempty" xml:"apptitle,omitempty"`
}
