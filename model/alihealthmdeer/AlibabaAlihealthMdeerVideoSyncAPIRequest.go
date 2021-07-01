package alihealthmdeer

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMdeerVideoSyncAPIRequest
合作伙伴视频同步给医知鹿接口 API请求
alibaba.alihealth.mdeer.video.sync

合伙做伴内容同步接口，用来视频同步 */
type AlibabaAlihealthMdeerVideoSyncAPIRequest struct {
	model.Params
	// 合作方头像url
	_partnerPortraitUrl string
	// 作者电话
	_phoneNumber string
	// 作者简介
	_authorIntroduction string
	// 作者科室
	_authorDepartment string
	// 作者级别
	_authorLevel string
	// 医院级别
	_hospitalLevel string
	// 医院名称
	_hospitalName string
	// 作者头像
	_portraitUrl string
	// 作者名称
	_authorName string
	// 作者id
	_authorId string
	// 合作方主页
	_partnerHomepage string
	// 合作方名称
	_partnerName string
	// 发布日期
	_releaseDate string
	// 视频文件url
	_videoFileUrl string
	// 视频落地页
	_videoMobileUrl string
	// 视频简介
	_videoIntroduction string
	// 视频长度
	_videoLength string
	// 视频所述疾病
	_disease string
	// 预览图url
	_priviewUrl string
	// 视频标题
	_videoTitle string
	// 视频id
	_videoId string
}

// NewAlibabaAlihealthMdeerVideoSyncRequest 初始化AlibabaAlihealthMdeerVideoSyncAPIRequest对象
func NewAlibabaAlihealthMdeerVideoSyncRequest() *AlibabaAlihealthMdeerVideoSyncAPIRequest {
	return &AlibabaAlihealthMdeerVideoSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.mdeer.video.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PartnerPortraitUrl Setter
// 合作方头像url
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPartnerPortraitUrl(_partnerPortraitUrl string) error {
	r._partnerPortraitUrl = _partnerPortraitUrl
	r.Set("partner_portrait_url", _partnerPortraitUrl)
	return nil
}

// Get PartnerPortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPartnerPortraitUrl() string {
	return r._partnerPortraitUrl
}

// Set is PhoneNumber Setter
// 作者电话
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPhoneNumber(_phoneNumber string) error {
	r._phoneNumber = _phoneNumber
	r.Set("phone_number", _phoneNumber)
	return nil
}

// Get PhoneNumber Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPhoneNumber() string {
	return r._phoneNumber
}

// Set is AuthorIntroduction Setter
// 作者简介
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetAuthorIntroduction(_authorIntroduction string) error {
	r._authorIntroduction = _authorIntroduction
	r.Set("author_introduction", _authorIntroduction)
	return nil
}

// Get AuthorIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetAuthorIntroduction() string {
	return r._authorIntroduction
}

// Set is AuthorDepartment Setter
// 作者科室
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetAuthorDepartment(_authorDepartment string) error {
	r._authorDepartment = _authorDepartment
	r.Set("author_department", _authorDepartment)
	return nil
}

// Get AuthorDepartment Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetAuthorDepartment() string {
	return r._authorDepartment
}

// Set is AuthorLevel Setter
// 作者级别
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetAuthorLevel(_authorLevel string) error {
	r._authorLevel = _authorLevel
	r.Set("author_level", _authorLevel)
	return nil
}

// Get AuthorLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetAuthorLevel() string {
	return r._authorLevel
}

// Set is HospitalLevel Setter
// 医院级别
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetHospitalLevel(_hospitalLevel string) error {
	r._hospitalLevel = _hospitalLevel
	r.Set("hospital_level", _hospitalLevel)
	return nil
}

// Get HospitalLevel Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetHospitalLevel() string {
	return r._hospitalLevel
}

// Set is HospitalName Setter
// 医院名称
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetHospitalName(_hospitalName string) error {
	r._hospitalName = _hospitalName
	r.Set("hospital_name", _hospitalName)
	return nil
}

// Get HospitalName Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetHospitalName() string {
	return r._hospitalName
}

// Set is PortraitUrl Setter
// 作者头像
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPortraitUrl(_portraitUrl string) error {
	r._portraitUrl = _portraitUrl
	r.Set("portrait_url", _portraitUrl)
	return nil
}

// Get PortraitUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPortraitUrl() string {
	return r._portraitUrl
}

// Set is AuthorName Setter
// 作者名称
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetAuthorName(_authorName string) error {
	r._authorName = _authorName
	r.Set("author_name", _authorName)
	return nil
}

// Get AuthorName Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetAuthorName() string {
	return r._authorName
}

// Set is AuthorId Setter
// 作者id
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetAuthorId(_authorId string) error {
	r._authorId = _authorId
	r.Set("author_id", _authorId)
	return nil
}

// Get AuthorId Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetAuthorId() string {
	return r._authorId
}

// Set is PartnerHomepage Setter
// 合作方主页
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPartnerHomepage(_partnerHomepage string) error {
	r._partnerHomepage = _partnerHomepage
	r.Set("partner_homepage", _partnerHomepage)
	return nil
}

// Get PartnerHomepage Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPartnerHomepage() string {
	return r._partnerHomepage
}

// Set is PartnerName Setter
// 合作方名称
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPartnerName(_partnerName string) error {
	r._partnerName = _partnerName
	r.Set("partner_name", _partnerName)
	return nil
}

// Get PartnerName Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPartnerName() string {
	return r._partnerName
}

// Set is ReleaseDate Setter
// 发布日期
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetReleaseDate(_releaseDate string) error {
	r._releaseDate = _releaseDate
	r.Set("release_date", _releaseDate)
	return nil
}

// Get ReleaseDate Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetReleaseDate() string {
	return r._releaseDate
}

// Set is VideoFileUrl Setter
// 视频文件url
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoFileUrl(_videoFileUrl string) error {
	r._videoFileUrl = _videoFileUrl
	r.Set("video_file_url", _videoFileUrl)
	return nil
}

// Get VideoFileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoFileUrl() string {
	return r._videoFileUrl
}

// Set is VideoMobileUrl Setter
// 视频落地页
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoMobileUrl(_videoMobileUrl string) error {
	r._videoMobileUrl = _videoMobileUrl
	r.Set("video_mobile_url", _videoMobileUrl)
	return nil
}

// Get VideoMobileUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoMobileUrl() string {
	return r._videoMobileUrl
}

// Set is VideoIntroduction Setter
// 视频简介
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoIntroduction(_videoIntroduction string) error {
	r._videoIntroduction = _videoIntroduction
	r.Set("video_introduction", _videoIntroduction)
	return nil
}

// Get VideoIntroduction Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoIntroduction() string {
	return r._videoIntroduction
}

// Set is VideoLength Setter
// 视频长度
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoLength(_videoLength string) error {
	r._videoLength = _videoLength
	r.Set("video_length", _videoLength)
	return nil
}

// Get VideoLength Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoLength() string {
	return r._videoLength
}

// Set is Disease Setter
// 视频所述疾病
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetDisease(_disease string) error {
	r._disease = _disease
	r.Set("disease", _disease)
	return nil
}

// Get Disease Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetDisease() string {
	return r._disease
}

// Set is PriviewUrl Setter
// 预览图url
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetPriviewUrl(_priviewUrl string) error {
	r._priviewUrl = _priviewUrl
	r.Set("priview_url", _priviewUrl)
	return nil
}

// Get PriviewUrl Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetPriviewUrl() string {
	return r._priviewUrl
}

// Set is VideoTitle Setter
// 视频标题
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoTitle(_videoTitle string) error {
	r._videoTitle = _videoTitle
	r.Set("video_title", _videoTitle)
	return nil
}

// Get VideoTitle Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoTitle() string {
	return r._videoTitle
}

// Set is VideoId Setter
// 视频id
func (r *AlibabaAlihealthMdeerVideoSyncAPIRequest) SetVideoId(_videoId string) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// Get VideoId Getter
func (r AlibabaAlihealthMdeerVideoSyncAPIRequest) GetVideoId() string {
	return r._videoId
}
