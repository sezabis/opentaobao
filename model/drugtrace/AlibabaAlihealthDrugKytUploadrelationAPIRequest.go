package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadrelationAPIRequest 关联关系上传 API请求
// alibaba.alihealth.drug.kyt.uploadrelation
//
// 关联关系上传
type AlibabaAlihealthDrugKytUploadrelationAPIRequest struct {
	model.Params
	// affirmFlag
	_affirmFlag string
	// fileContent(可不添)
	_fileContent string
	// 加密之后的文件内容字符串
	_fileContentString string
	// 上传文件的企业ID
	_refEntId string
	// 客户端类型
	_clientType string
	// 关联关系文件信息
	_saveCodeRelation *SaveCodeRelationType
}

// NewAlibabaAlihealthDrugKytUploadrelationRequest 初始化AlibabaAlihealthDrugKytUploadrelationAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadrelationRequest() *AlibabaAlihealthDrugKytUploadrelationAPIRequest {
	return &AlibabaAlihealthDrugKytUploadrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.uploadrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAffirmFlag is AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetAffirmFlag(_affirmFlag string) error {
	r._affirmFlag = _affirmFlag
	r.Set("affirm_flag", _affirmFlag)
	return nil
}

// GetAffirmFlag AffirmFlag Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetAffirmFlag() string {
	return r._affirmFlag
}

// SetFileContent is FileContent Setter
// fileContent(可不添)
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetFileContentString is FileContentString Setter
// 加密之后的文件内容字符串
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetFileContentString(_fileContentString string) error {
	r._fileContentString = _fileContentString
	r.Set("file_content_string", _fileContentString)
	return nil
}

// GetFileContentString FileContentString Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetFileContentString() string {
	return r._fileContentString
}

// SetRefEntId is RefEntId Setter
// 上传文件的企业ID
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetClientType() string {
	return r._clientType
}

// SetSaveCodeRelation is SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
	r._saveCodeRelation = _saveCodeRelation
	r.Set("save_code_relation", _saveCodeRelation)
	return nil
}

// GetSaveCodeRelation SaveCodeRelation Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetSaveCodeRelation() *SaveCodeRelationType {
	return r._saveCodeRelation
}
