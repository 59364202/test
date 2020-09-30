package public

import (
	"encoding/json"
	"time"

	"haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"

	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_drought_area "haii.or.th/api/thaiwater30/model/drought_area"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_rainforecast "haii.or.th/api/thaiwater30/model/rainforecast"
	model_storm "haii.or.th/api/thaiwater30/model/storm"
	model_swan "haii.or.th/api/thaiwater30/model/swan"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
)

type ThailandStruct struct {
	Date             *time.Time            `json:"date,omitempty"`
	Dam              *ThailandModuleStruct `json:"dam,omitempty"`
	PreRain          *ThailandModuleStruct `json:"pre_rain,omitempty"`
	PreRainSea       *ThailandModuleStruct `json:"pre_rain_sea,omitempty"`
	PreRainAsia      *ThailandModuleStruct `json:"pre_rain_asia,omitempty"`
	PreRainBasin     *ThailandModuleStruct `json:"pre_rain_basin,omitempty"`
	PreRainAnimation *result.Result        `json:"pre_rain_animation,omitempty"`
	Province         *result.Result        `json:"province,omitempty"`
	Radar            *ThailandModuleStruct `json:"radar,omitempty"`
	Rain             *ThailandModuleStruct `json:"rain,omitempty"`
	Storm            *ThailandModuleStruct `json:"storm,omitempty"`
	Warning          *ThailandModuleStruct `json:"warning,omitempty"`
	WaterLevel       *ThailandModuleStruct `json:"waterlevel,omitempty"`
	WaterQuality     *ThailandModuleStruct `json:"waterquality,omitempty"`
	Wave             *ThailandModuleStruct `json:"wave,omitempty"`
	WaveAnimation    *result.Result        `json:"wave_animation,omitempty"`
}

type ThaiwaterMainStruct struct {
	Date        *time.Time            `json:"date,omitempty"`
	Rain        *ThailandModuleStruct `json:"rain,omitempty"`
	Temperature *ThailandModuleStruct `json:"temperature,omitempty"`
	Dam         *ThailandModuleStruct `json:"dam,omitempty"`
	Storm       *ThailandModuleStruct `json:"storm,omitempty"`
}

type ThaiwaterWeatherStruct struct {
	Rain        *result.Result `json:"rain,omitempty"`
	Temperature *result.Result `json:"temperature,omitempty"`
	Humid       *result.Result `json:"humid,omitempty"`
	Pressure    *result.Result `json:"pressure,omitempty"`
	Wind        *result.Result `json:"wind,omitempty"`
}

type ThailandModuleStruct struct {
	Data        *result.Result  `json:"data"`
	Setting     json.RawMessage `json:"setting,omitempty"`
	Date        *time.Time      `json:"date,omitempty"`
	TempData    *result.Result  `json:"temp_data,omitempty"`
	TempData2   *result.Result  `json:"temp_data2,omitempty"`
	Drought     *result.Result  `json:"drought,omitempty"`
	Flood       *result.Result  `json:"flood,omitempty"`
	Summary     *result.Result  `json:"summary,omitempty"`
	Summary4Dam *result.Result  `json:"summary4dam,omitempty"`
}

// province
type Struct_Province struct {
	Result string                          `json:"result"` // example:`OK`
	Data   []*model_geocode.Struct_Geocode `json:"data"`   // ข้อมูล
}

// rain
type Struct_Rain struct {
	Data    *Struct_Rain_Data            `json:"data"`    // ฝน
	Setting *uSetting.Struct_RainSetting `json:"setting"` // เกณฑ์
}
type Struct_Rain_Data struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_rainfall24hr.Struct_Rainfall24H `json:"data"`   // ฝน
}

// radar
type Struct_Radar struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Radar `json:"data"`   // ภาพเรเดาร์
}

// waterlevel
type Struct_WaterLevel struct {
	Data    *Struct_WaterLevel_Data            `json:"data"`    // ระดับน้ำ
	Setting *uSetting.Struct_WaterlevelSetting `json:"setting"` // เกณฑ์
}
type Struct_WaterLevel_Data struct {
	Result string                                     `json:"result"` // example:`OK`
	Data   []*model_tele_waterlevel.Struct_Waterlevel `json:"data"`   // ระดับน้ำ
}

// dam
type Struct_Dam struct {
	Data    *Struct_Dam_Data               `json:"data"`    // เขื่อน
	Setting *uSetting.Struct_DamScaleColor `json:"setting"` // เกณฑ์
}
type Struct_Dam_Data struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_dam_daily.Struct_DamDaily `json:"data"`   // เขื่อน
}

// waterquality
type Struct_WaterQuality struct {
	Data    *Struct_WaterQuality_Data            `json:"data"`    // คุณภาพน้ำ
	Setting *uSetting.Struct_WaterqualitySetting `json:"setting"` // เกณฑ์
}
type Struct_WaterQuality_Data struct {
	Result string                                    `json:"result"` // example:`OK`
	Data   []*model_waterquality.Struct_WaterQuality `json:"data"`   // คุุณภาพน้ำ
}

// storm
type Struct_St struct {
	Data    *Struct_Storm                   `json:"data"`    // ภาพพายุ
	Setting []*uSetting.Struct_StormSetting `json:"setting"` // เกณฑ์
}

type Struct_Storm struct {
	Result string             `json:"result"` // example:`OK`
	Data   *Struct_Storm_Data `json:"data"`   // ภาพพายุ
}
type Struct_Storm_Data struct {
	C []*model_latest_media.Struct_Media `json:"college"` //  University College London
	K []*model_latest_media.Struct_Media `json:"kochi"`   //  kochi
	T []*model_latest_media.Struct_Media `json:"typhoon"` //  Digital Typhoon
	U []*model_latest_media.Struct_Media `json:"us"`      //  US Naval Research Laboratory
}
type Struct_Md struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Media `json:"data"`   // ข้อมูลสื่อ
}
type Index_Storm struct {
	Data    *Index_Storm_Data               `json:"data"`    // คาดการณ์พายุ
	Setting []*uSetting.Struct_StormSetting `json:"setting"` // เกณฑ์
}

type Index_Storm_Data struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model_storm.Struct_Strom `json:"data"`   // พายุ
}

// warning
type Struct_Warning struct {
	Setting *uSetting.Struct_WarningSetting `json:"setting"`    // เกณฑ์
	Drought *Struct_Warning_Drought         `json:"drought"`    // พื้นที่ประสบภัยแล้ง
	Flood   *Struct_Warning_Flood           `json:"flood"`      // พื้นที่แจ้งเตือนอุทกภัยจาก ศภช.
	TempD   *Struct_Rain_Data               `json:"temp_data"`  // ข้อมูลฝนจากหน้าแผนที่ประเทศไทย
	TempD2  *Index_PreRain_Data             `json:"temp_data2"` // ข้อมูลคาดการณ์ฝนจากหน้าแผนที่ประเทศไทย
}
type Struct_Warning_Drought struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_drought_area.Struct_DroughtArea `json:"data'`   // พื้นที่ประสบภัยแล้ง
}
type Struct_Warning_Flood struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_drought_area.Struct_DroughtArea `json:"data'`   // พื้นที่แจ้งเตือนอุทกภัยจาก ศภช.
}

// prerain
type Index_PreRain struct {
	Data    *Index_PreRain_Data             `json:"data"`    // คาดการณ์ฝน
	Setting *uSetting.Struct_PreRainSetting `json:"setting"` // เกณฑ์
}
type Index_PreRain_Data struct {
	Result string                                    `json:"result"` // example:`OK`
	Data   []*model_rainforecast.Struct_Rainforecast `json:"data"`   // คาดการณ์ฝน
}

// wave
type Index_Wave struct {
	Data    *Index_Wave_Data             `json:"data"`    // คาดการณ์คลื่น
	Setting *uSetting.Struct_WaveSetting `json:"setting"` // เกณฑ์
}
type Index_Wave_Data struct {
	Result string                    `json:"result"` // example:`OK`
	Data   []*model_swan.Struct_Swan `json:"data"`   // คาดการณ์คลื่น
}

// ----------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------- thailand main struct  --------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------
type Struct_Main struct {
	Dam              *Struct_Dam          `json:"dam"`                // เขื่อน
	PreRain          *Struct_Md           `json:"pre_rain"`           // ภาพคาดการณ์ฝนล่วงหน้า 7 วัน
	PreRainAnimation *Struct_Md           `json:"pre_rain_animation"` // animation คาดการณ์ฝน
	Province         *Struct_Province     `json:"province"`           // รายชื่อจังหวัด
	Radar            *Struct_Radar        `json:"radar"`              // ภาพเรเดาร์
	Rain             *Struct_Rain         `json:"rain"`               // ฝนย้อนหลัง 24 ชม
	Storm            *Struct_St           `json:"storm"`              // ภาพพายุ
	Warning          *Struct_Warning      `json:"warning,omitempty"`  // พื้นที่ประกาศภัย
	WaterLevel       *Struct_WaterLevel   `json:"waterlevel"`         // ระดับน้ำ
	WaterQuality     *Struct_WaterQuality `json:"waterquality"`       // คุณภาพน้ำ
	Wave             *Struct_Md           `json:"wave"`               // ภาพคาดการณ์คลื่นล่วงหน้า 7 วัน
	WaveAnimation    *Struct_Md           `json:"wave_animation"`     // animation คาดการณ์คลื่น
}

// ----------------------------------------------------------------------------------------------------------------------
// --------------------------------------------------- thailand index struct  -------------------------------------------
// ----------------------------------------------------------------------------------------------------------------------
type Index struct {
	Dam          *Struct_Dam          `json:"dam"`          // เขื่อน
	PreRain      *Index_PreRain       `json:"pre_rain"`     // คาดการณ์ฝน
	Province     *Struct_Province     `json:"province"`     // รายชื่อจังหวัด
	Rain         *Struct_Rain         `json:"rain"`         // ฝน
	Storm        *Index_Storm         `json:"storm"`        // พายุ
	Warning      *Struct_Warning      `json:"warning"`      // พื้นที่ประกาศภัย
	WaterLevel   *Struct_WaterLevel   `json:"waterlevel"`   // ระดับน้ำ
	WaterQuality *Struct_WaterQuality `json:"waterquality"` // คุณภาพน้ำ
	Wave         *Index_Wave          `json:"wave"`         // คาดการณ์คลื่น
}
