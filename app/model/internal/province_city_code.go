// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// ProvinceCityCode is the golang structure for table province_city_code.
type ProvinceCityCode struct {
	Code         string `orm:"code"          json:"code"`         // 六位地区码
	Name         string `orm:"name"          json:"name"`         // 六位地区码含义
	ProvinceCode string `orm:"province_code" json:"provinceCode"` // 前两位省市码
	ProvinceName string `orm:"province_name" json:"provinceName"` // 前两位省市码含义
	CityCode     string `orm:"city_code"     json:"cityCode"`     // 前四位市县码
	CityName     string `orm:"city_name"     json:"cityName"`     // 前四位市县码含义
	AreaCode     string `orm:"area_code"     json:"areaCode"`     // 前六位地区码
	AreaName     string `orm:"area_name"     json:"areaName"`     // 前六位地区码含义
}
