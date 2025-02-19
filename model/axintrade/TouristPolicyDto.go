package axintrade

// TouristPolicyDto 结构体
type TouristPolicyDto struct {
	// 游客信息填写字段
	FieldTypes []string `json:"field_types,omitempty" xml:"field_types>string,omitempty"`
	// 证件类型
	CertificatesTypes []string `json:"certificates_types,omitempty" xml:"certificates_types>string,omitempty"`
}
