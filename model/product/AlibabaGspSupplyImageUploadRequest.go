package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gsp图片上传 API请求
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间
*/
type AlibabaGspSupplyImageUploadRequest struct {
    model.Params
    // 图片名称
    _fileName   string
    // 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
    _fileContent   *model.File
}

// 初始化AlibabaGspSupplyImageUploadRequest对象
func NewAlibabaGspSupplyImageUploadRequest() *AlibabaGspSupplyImageUploadRequest{
    return &AlibabaGspSupplyImageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGspSupplyImageUploadRequest) GetApiMethodName() string {
    return "alibaba.gsp.supply.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGspSupplyImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileName Setter
// 图片名称
func (r *AlibabaGspSupplyImageUploadRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaGspSupplyImageUploadRequest) GetFileName() string {
    return r._fileName
}
// FileContent Setter
// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
func (r *AlibabaGspSupplyImageUploadRequest) SetFileContent(_fileContent *model.File) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaGspSupplyImageUploadRequest) GetFileContent() *model.File {
    return r._fileContent
}
