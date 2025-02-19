package dt

// TaobaoCmedicalFaceDetectionCallbackResult 结构体
type TaobaoCmedicalFaceDetectionCallbackResult struct {
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功，true：成功，false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
