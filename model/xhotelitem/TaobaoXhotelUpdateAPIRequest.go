package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelUpdateAPIRequest 酒店更新接口（ID不存在自动新增） API请求
// taobao.xhotel.update
//
// 酒店更新接口
type TaobaoXhotelUpdateAPIRequest struct {
	model.Params
	// 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
	_name string
	// 酒店曾用名
	_usedName string
	// domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
	_country string
	// 商业区（圈）长度不超过20字
	_business string
	// 酒店地址。长度不能超过255
	_address string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
	_positionType string
	// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 不要使用
	_extend string
	// 必传，酒店标识，商家酒店ID
	_outerId string
	// 系统商，一般情况不用，需申请使用
	_vendor string
	// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
	_star string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 装修时间，格式为2015-10-01
	_decorateTime string
	// 楼层信息
	_floors string
	// 酒店描述
	_description string
	// 酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
	_hotelFacilities string
	// 酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
	_service string
	// 房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
	_roomFacilities string
	// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
	_pics string
	// 酒店品牌。取值为数字。枚举见链接：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.S16vXH&docType=1&articleId=120180
	_brand string
	// 邮编
	_postalCode string
	// 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}
	_hotelPolicies string
	// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
	_bookingNotice string
	// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
	_creditCardTypes string
	// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
	_orbitTrack string
	// 卖家酒店英文名称
	_nameE string
	// 打标去标使用的 tagJson 字段
	_tagJson string
	// 旺旺昵称
	_aliNick string
	// 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
	_supplier string
	// 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardBookingNotice string
	// 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardHotelFacilities string
	// 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardHotelService string
	// 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardRoomFacilities string
	// 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardAmuseFacilities string
	// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
	_coordinateSystem string
	// （已废弃）请使用outer_id来标识要修改的酒店
	_hid int64
	// 是否国内酒店。0:国内;1:国外
	_domestic int64
	// 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
	_province int64
	// 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
	_city int64
	// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
	_district int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
	_shid int64
	// 房间数 0~9999之内的数字
	_rooms int64
	// 酒店状态 0:正常，-1:删除，-2:停售
	_status *model.File
	// 0:酒店；1:客栈
	_hotelType int64
	// 0:可以接待外宾；1:仅内宾
	_serviceType int64
}

// NewTaobaoXhotelUpdateRequest 初始化TaobaoXhotelUpdateAPIRequest对象
func NewTaobaoXhotelUpdateRequest() *TaobaoXhotelUpdateAPIRequest {
	return &TaobaoXhotelUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetName is Name Setter
// 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
func (r *TaobaoXhotelUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelUpdateAPIRequest) GetName() string {
	return r._name
}

// SetUsedName is UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelUpdateAPIRequest) SetUsedName(_usedName string) error {
	r._usedName = _usedName
	r.Set("used_name", _usedName)
	return nil
}

// GetUsedName UsedName Getter
func (r TaobaoXhotelUpdateAPIRequest) GetUsedName() string {
	return r._usedName
}

// SetCountry is Country Setter
// domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
func (r *TaobaoXhotelUpdateAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r TaobaoXhotelUpdateAPIRequest) GetCountry() string {
	return r._country
}

// SetBusiness is Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelUpdateAPIRequest) SetBusiness(_business string) error {
	r._business = _business
	r.Set("business", _business)
	return nil
}

// GetBusiness Business Getter
func (r TaobaoXhotelUpdateAPIRequest) GetBusiness() string {
	return r._business
}

// SetAddress is Address Setter
// 酒店地址。长度不能超过255
func (r *TaobaoXhotelUpdateAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoXhotelUpdateAPIRequest) GetAddress() string {
	return r._address
}

// SetLongitude is Longitude Setter
// 经度
func (r *TaobaoXhotelUpdateAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoXhotelUpdateAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *TaobaoXhotelUpdateAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoXhotelUpdateAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetPositionType is PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelUpdateAPIRequest) SetPositionType(_positionType string) error {
	r._positionType = _positionType
	r.Set("position_type", _positionType)
	return nil
}

// GetPositionType PositionType Getter
func (r TaobaoXhotelUpdateAPIRequest) GetPositionType() string {
	return r._positionType
}

// SetTel is Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelUpdateAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// GetTel Tel Getter
func (r TaobaoXhotelUpdateAPIRequest) GetTel() string {
	return r._tel
}

// SetExtend is Extend Setter
// 不要使用
func (r *TaobaoXhotelUpdateAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoXhotelUpdateAPIRequest) GetExtend() string {
	return r._extend
}

// SetOuterId is OuterId Setter
// 必传，酒店标识，商家酒店ID
func (r *TaobaoXhotelUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，一般情况不用，需申请使用
func (r *TaobaoXhotelUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetStar is Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelUpdateAPIRequest) SetStar(_star string) error {
	r._star = _star
	r.Set("star", _star)
	return nil
}

// GetStar Star Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStar() string {
	return r._star
}

// SetOpeningTime is OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelUpdateAPIRequest) SetOpeningTime(_openingTime string) error {
	r._openingTime = _openingTime
	r.Set("opening_time", _openingTime)
	return nil
}

// GetOpeningTime OpeningTime Getter
func (r TaobaoXhotelUpdateAPIRequest) GetOpeningTime() string {
	return r._openingTime
}

// SetDecorateTime is DecorateTime Setter
// 装修时间，格式为2015-10-01
func (r *TaobaoXhotelUpdateAPIRequest) SetDecorateTime(_decorateTime string) error {
	r._decorateTime = _decorateTime
	r.Set("decorate_time", _decorateTime)
	return nil
}

// GetDecorateTime DecorateTime Getter
func (r TaobaoXhotelUpdateAPIRequest) GetDecorateTime() string {
	return r._decorateTime
}

// SetFloors is Floors Setter
// 楼层信息
func (r *TaobaoXhotelUpdateAPIRequest) SetFloors(_floors string) error {
	r._floors = _floors
	r.Set("floors", _floors)
	return nil
}

// GetFloors Floors Getter
func (r TaobaoXhotelUpdateAPIRequest) GetFloors() string {
	return r._floors
}

// SetDescription is Description Setter
// 酒店描述
func (r *TaobaoXhotelUpdateAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoXhotelUpdateAPIRequest) GetDescription() string {
	return r._description
}

// SetHotelFacilities is HotelFacilities Setter
// 酒店设施。json格式示例值：{&#34;free Wi-Fi in all rooms&#34;:&#34;true&#34;,&#34;massage&#34;:&#34;true&#34;,&#34;meetingRoom&#34;:&#34;true&#34;}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
func (r *TaobaoXhotelUpdateAPIRequest) SetHotelFacilities(_hotelFacilities string) error {
	r._hotelFacilities = _hotelFacilities
	r.Set("hotel_facilities", _hotelFacilities)
	return nil
}

// GetHotelFacilities HotelFacilities Getter
func (r TaobaoXhotelUpdateAPIRequest) GetHotelFacilities() string {
	return r._hotelFacilities
}

// SetService is Service Setter
// 酒店基础服务。json格式示例值：{&#34;receiveForeignGuests&#34;:&#34;true&#34;,&#34;morningCall&#34;:&#34;true&#34;,&#34;breakfast&#34;:&#34;true&#34;}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
func (r *TaobaoXhotelUpdateAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoXhotelUpdateAPIRequest) GetService() string {
	return r._service
}

// SetRoomFacilities is RoomFacilities Setter
// 房间的基础设施。json格式示例值：{&#34;bathtub&#34;:&#34;true&#34;,&#34;bathPub&#34;:&#34;true&#34;}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
func (r *TaobaoXhotelUpdateAPIRequest) SetRoomFacilities(_roomFacilities string) error {
	r._roomFacilities = _roomFacilities
	r.Set("room_facilities", _roomFacilities)
	return nil
}

// GetRoomFacilities RoomFacilities Getter
func (r TaobaoXhotelUpdateAPIRequest) GetRoomFacilities() string {
	return r._roomFacilities
}

// SetPics is Pics Setter
// 酒店图片只支持远程图片，格式如下：[{&#34;url&#34;:&#34;http://123.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;,&#34;type&#34;:&#34;大堂&#34;,&#34;attribute&#34;:&#34;普通图&#34;},{&#34;url&#34;:&#34;http://456.jpg&#34;,&#34;ismain&#34;:&#34;true&#34;,&#34;type&#34;:&#34;公共区域&#34;,&#34;attribute&#34;:&#34;全景图&#34;},{&#34;url&#34;:&#34;http://789.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;,&#34;type&#34;:&#34;大堂&#34;,&#34;attribute&#34;:&#34;普通图&#34;}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
func (r *TaobaoXhotelUpdateAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoXhotelUpdateAPIRequest) GetPics() string {
	return r._pics
}

// SetBrand is Brand Setter
// 酒店品牌。取值为数字。枚举见链接：https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.S16vXH&amp;docType=1&amp;articleId=120180
func (r *TaobaoXhotelUpdateAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaoXhotelUpdateAPIRequest) GetBrand() string {
	return r._brand
}

// SetPostalCode is PostalCode Setter
// 邮编
func (r *TaobaoXhotelUpdateAPIRequest) SetPostalCode(_postalCode string) error {
	r._postalCode = _postalCode
	r.Set("postal_code", _postalCode)
	return nil
}

// GetPostalCode PostalCode Getter
func (r TaobaoXhotelUpdateAPIRequest) GetPostalCode() string {
	return r._postalCode
}

// SetHotelPolicies is HotelPolicies Setter
// 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{&#34;children_age_from&#34;:&#34;2&#34;,&#34;children_age_to&#34;:&#34;3&#34;,&#34;children_stay_free&#34;:&#34;True&#34;,&#34;infant_age&#34;:&#34;1&#34;,&#34;min_guest_age&#34;:&#34;4&#34;}
func (r *TaobaoXhotelUpdateAPIRequest) SetHotelPolicies(_hotelPolicies string) error {
	r._hotelPolicies = _hotelPolicies
	r.Set("hotel_policies", _hotelPolicies)
	return nil
}

// GetHotelPolicies HotelPolicies Getter
func (r TaobaoXhotelUpdateAPIRequest) GetHotelPolicies() string {
	return r._hotelPolicies
}

// SetBookingNotice is BookingNotice Setter
// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
func (r *TaobaoXhotelUpdateAPIRequest) SetBookingNotice(_bookingNotice string) error {
	r._bookingNotice = _bookingNotice
	r.Set("booking_notice", _bookingNotice)
	return nil
}

// GetBookingNotice BookingNotice Getter
func (r TaobaoXhotelUpdateAPIRequest) GetBookingNotice() string {
	return r._bookingNotice
}

// SetCreditCardTypes is CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelUpdateAPIRequest) SetCreditCardTypes(_creditCardTypes string) error {
	r._creditCardTypes = _creditCardTypes
	r.Set("credit_card_types", _creditCardTypes)
	return nil
}

// GetCreditCardTypes CreditCardTypes Getter
func (r TaobaoXhotelUpdateAPIRequest) GetCreditCardTypes() string {
	return r._creditCardTypes
}

// SetOrbitTrack is OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelUpdateAPIRequest) SetOrbitTrack(_orbitTrack string) error {
	r._orbitTrack = _orbitTrack
	r.Set("orbit_track", _orbitTrack)
	return nil
}

// GetOrbitTrack OrbitTrack Getter
func (r TaobaoXhotelUpdateAPIRequest) GetOrbitTrack() string {
	return r._orbitTrack
}

// SetNameE is NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelUpdateAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoXhotelUpdateAPIRequest) GetNameE() string {
	return r._nameE
}

// SetTagJson is TagJson Setter
// 打标去标使用的 tagJson 字段
func (r *TaobaoXhotelUpdateAPIRequest) SetTagJson(_tagJson string) error {
	r._tagJson = _tagJson
	r.Set("tag_json", _tagJson)
	return nil
}

// GetTagJson TagJson Getter
func (r TaobaoXhotelUpdateAPIRequest) GetTagJson() string {
	return r._tagJson
}

// SetAliNick is AliNick Setter
// 旺旺昵称
func (r *TaobaoXhotelUpdateAPIRequest) SetAliNick(_aliNick string) error {
	r._aliNick = _aliNick
	r.Set("ali_nick", _aliNick)
	return nil
}

// GetAliNick AliNick Getter
func (r TaobaoXhotelUpdateAPIRequest) GetAliNick() string {
	return r._aliNick
}

// SetSupplier is Supplier Setter
// 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
func (r *TaobaoXhotelUpdateAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelUpdateAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetSettlementCurrency is SettlementCurrency Setter
// 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelUpdateAPIRequest) SetSettlementCurrency(_settlementCurrency string) error {
	r._settlementCurrency = _settlementCurrency
	r.Set("settlement_currency", _settlementCurrency)
	return nil
}

// GetSettlementCurrency SettlementCurrency Getter
func (r TaobaoXhotelUpdateAPIRequest) GetSettlementCurrency() string {
	return r._settlementCurrency
}

// SetStandardBookingNotice is StandardBookingNotice Setter
// 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=108891
func (r *TaobaoXhotelUpdateAPIRequest) SetStandardBookingNotice(_standardBookingNotice string) error {
	r._standardBookingNotice = _standardBookingNotice
	r.Set("standard_booking_notice", _standardBookingNotice)
	return nil
}

// GetStandardBookingNotice StandardBookingNotice Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStandardBookingNotice() string {
	return r._standardBookingNotice
}

// SetStandardHotelFacilities is StandardHotelFacilities Setter
// 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=108891
func (r *TaobaoXhotelUpdateAPIRequest) SetStandardHotelFacilities(_standardHotelFacilities string) error {
	r._standardHotelFacilities = _standardHotelFacilities
	r.Set("standard_hotel_facilities", _standardHotelFacilities)
	return nil
}

// GetStandardHotelFacilities StandardHotelFacilities Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStandardHotelFacilities() string {
	return r._standardHotelFacilities
}

// SetStandardHotelService is StandardHotelService Setter
// 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=108891
func (r *TaobaoXhotelUpdateAPIRequest) SetStandardHotelService(_standardHotelService string) error {
	r._standardHotelService = _standardHotelService
	r.Set("standard_hotel_service", _standardHotelService)
	return nil
}

// GetStandardHotelService StandardHotelService Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStandardHotelService() string {
	return r._standardHotelService
}

// SetStandardRoomFacilities is StandardRoomFacilities Setter
// 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=108891
func (r *TaobaoXhotelUpdateAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// GetStandardRoomFacilities StandardRoomFacilities Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}

// SetStandardAmuseFacilities is StandardAmuseFacilities Setter
// 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=108891
func (r *TaobaoXhotelUpdateAPIRequest) SetStandardAmuseFacilities(_standardAmuseFacilities string) error {
	r._standardAmuseFacilities = _standardAmuseFacilities
	r.Set("standard_amuse_facilities", _standardAmuseFacilities)
	return nil
}

// GetStandardAmuseFacilities StandardAmuseFacilities Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStandardAmuseFacilities() string {
	return r._standardAmuseFacilities
}

// SetCoordinateSystem is CoordinateSystem Setter
// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
func (r *TaobaoXhotelUpdateAPIRequest) SetCoordinateSystem(_coordinateSystem string) error {
	r._coordinateSystem = _coordinateSystem
	r.Set("coordinate_system", _coordinateSystem)
	return nil
}

// GetCoordinateSystem CoordinateSystem Getter
func (r TaobaoXhotelUpdateAPIRequest) GetCoordinateSystem() string {
	return r._coordinateSystem
}

// SetHid is Hid Setter
// （已废弃）请使用outer_id来标识要修改的酒店
func (r *TaobaoXhotelUpdateAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelUpdateAPIRequest) GetHid() int64 {
	return r._hid
}

// SetDomestic is Domestic Setter
// 是否国内酒店。0:国内;1:国外
func (r *TaobaoXhotelUpdateAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r TaobaoXhotelUpdateAPIRequest) GetDomestic() int64 {
	return r._domestic
}

// SetProvince is Province Setter
// 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
func (r *TaobaoXhotelUpdateAPIRequest) SetProvince(_province int64) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaoXhotelUpdateAPIRequest) GetProvince() int64 {
	return r._province
}

// SetCity is City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
func (r *TaobaoXhotelUpdateAPIRequest) SetCity(_city int64) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoXhotelUpdateAPIRequest) GetCity() int64 {
	return r._city
}

// SetDistrict is District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
func (r *TaobaoXhotelUpdateAPIRequest) SetDistrict(_district int64) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// GetDistrict District Getter
func (r TaobaoXhotelUpdateAPIRequest) GetDistrict() int64 {
	return r._district
}

// SetShid is Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelUpdateAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoXhotelUpdateAPIRequest) GetShid() int64 {
	return r._shid
}

// SetRooms is Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelUpdateAPIRequest) SetRooms(_rooms int64) error {
	r._rooms = _rooms
	r.Set("rooms", _rooms)
	return nil
}

// GetRooms Rooms Getter
func (r TaobaoXhotelUpdateAPIRequest) GetRooms() int64 {
	return r._rooms
}

// SetStatus is Status Setter
// 酒店状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelUpdateAPIRequest) SetStatus(_status *model.File) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelUpdateAPIRequest) GetStatus() *model.File {
	return r._status
}

// SetHotelType is HotelType Setter
// 0:酒店；1:客栈
func (r *TaobaoXhotelUpdateAPIRequest) SetHotelType(_hotelType int64) error {
	r._hotelType = _hotelType
	r.Set("hotel_type", _hotelType)
	return nil
}

// GetHotelType HotelType Getter
func (r TaobaoXhotelUpdateAPIRequest) GetHotelType() int64 {
	return r._hotelType
}

// SetServiceType is ServiceType Setter
// 0:可以接待外宾；1:仅内宾
func (r *TaobaoXhotelUpdateAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TaobaoXhotelUpdateAPIRequest) GetServiceType() int64 {
	return r._serviceType
}
