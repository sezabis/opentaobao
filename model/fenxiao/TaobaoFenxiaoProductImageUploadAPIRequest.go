package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductImageUploadAPIRequest
产品图片上传 API请求
taobao.fenxiao.product.image.upload

产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片 */
type TaobaoFenxiaoProductImageUploadAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 产品主图图片空间相对路径或绝对路径
	_picPath string
	// 产品图片
	_image *model.File
	// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
	_position int64
	// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
	_properties string
}

// NewTaobaoFenxiaoProductImageUploadRequest 初始化TaobaoFenxiaoProductImageUploadAPIRequest对象
func NewTaobaoFenxiaoProductImageUploadRequest() *TaobaoFenxiaoProductImageUploadAPIRequest {
	return &TaobaoFenxiaoProductImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductImageUploadAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetProductId() int64 {
	return r._productId
}

// Set is PicPath Setter
// 产品主图图片空间相对路径或绝对路径
func (r *TaobaoFenxiaoProductImageUploadAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// Get PicPath Getter
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetPicPath() string {
	return r._picPath
}

// Set is Image Setter
// 产品图片
func (r *TaobaoFenxiaoProductImageUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// Get Image Getter
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetImage() *model.File {
	return r._image
}

// Set is Position Setter
// 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图
func (r *TaobaoFenxiaoProductImageUploadAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// Get Position Getter
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetPosition() int64 {
	return r._position
}

// Set is Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaoFenxiaoProductImageUploadAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// Get Properties Getter
func (r TaobaoFenxiaoProductImageUploadAPIRequest) GetProperties() string {
	return r._properties
}
