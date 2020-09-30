package setting

import (
	"encoding/json"

	"haii.or.th/api/server/model/setting"
)

//
// Frontend.public.rain_setting
//
type Struct_RainSetting struct {
	Level *Struct_RainSetting_L   `json:"level_color"` // ระดับ
	Rule  *Struct_RainSetting_R   `json:"ruule"`       // เกณฑ์หน้าแผนที่ประเทศไทย
	Scale []*Struct_RainSetting_S `json:"scale"`       // เกณฑ์หน้าหลัก
}
type Struct_RainSetting_L struct {
	L3 *Struct_RainSetting_Level `json:"3"` // การแสดงผลตามระดับ
	L4 *Struct_RainSetting_Level `json:"4"` // การแสดงผลตามระดับ
}
type Struct_RainSetting_Level struct {
	Color     string `json:"color"`     // example:`#FE8A04` สี
	ColorName string `json:"colorname"` // example:`orange` ชื่อสี
	Name      string `json:"name"`      // example:`เฝ้าระวังพิเศษ` ชื่อระดับ
	Trans     string `json:"trans"`     // example:`rain_level_3` ชื่อ trans ใน laravel
}

type Struct_RainSetting_R struct {
	R1  []*Struct_RainSetting_Rule `json:"01"` // เกณฑ์ตามโซน
	R2  []*Struct_RainSetting_Rule `json:"02"` // เกณฑ์ตามโซน
	R3  []*Struct_RainSetting_Rule `json:"03"` // เกณฑ์ตามโซน
	R4  []*Struct_RainSetting_Rule `json:"04"` // เกณฑ์ตามโซน
	R5  []*Struct_RainSetting_Rule `json:"05"` // เกณฑ์ตามโซน
	R6  []*Struct_RainSetting_Rule `json:"06"` // เกณฑ์ตามโซน
	R7  []*Struct_RainSetting_Rule `json:"07"` // เกณฑ์ตามโซน
	R8  []*Struct_RainSetting_Rule `json:"08"` // เกณฑ์ตามโซน
	R9  []*Struct_RainSetting_Rule `json:"09"` // เกณฑ์ตามโซน
	R10 []*Struct_RainSetting_Rule `json:"10"` // เกณฑ์ตามโซน
	R11 []*Struct_RainSetting_Rule `json:"11"` // เกณฑ์ตามโซน
}
type Struct_RainSetting_Rule struct {
	Operator string `json:"operator"` // example:`>=` การกระทำ
	Rain1h   string `json:"rain1h"`   // example:`30` ฝน 1 ชม.
	Rain24h  string `json:"rain24h"`  // example:`80` ฝน 24 ชม.
	Rain3d   string `json:"rain3d"`   // example:`160` ฝน 3 วัน
	Level    int    `json:"level"`    // exmample:`4` ระดับ
}

type Struct_RainSetting_S struct {
	Operator       string `json:"operator"`       // example:`>` การกระทำ
	Term           string `json:"term"`           // example:`90` ค่าฝน
	Color          string `json:"color"`          // example:`#EE141F` สี
	ColorName      string `json:"colorname"`      // example:`red` ชื่อสี
	Criterion_text string `json:"criterion_text"` // example:`ฝนตกหนักมาก`
}

//	สร้างตัว Struct_RainSetting พร้อม setting ค่าจาก db
func New_Struct_RainSetting() (*Struct_RainSetting, error) {
	s := new(Struct_RainSetting)
	err := setting.GetSystemSettingPtr("Frontend.public.rain_setting", &s)
	if err != nil {
		return nil, err
	}
	return s, nil

}

//  เปรียบเทียบค่า v กับ struct Scale
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_RainSetting) CompareRain(v interface{}) *Struct_RainSetting_S {
	if s.Scale == nil {
		return nil
	}
	for _, sc := range s.Scale {
		if sc.Compare(v) {
			return sc
		}
	}
	return nil
}

//  เปรียบเทียบค่า v กับ struct
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_RainSetting_S) Compare(v interface{}) bool {
	b, err := Compare(v, s.Operator, s.Term)
	if err != nil {
		return false
	}
	return b
}

//
// Frontend.public.dam_scale_color
//
type Struct_DamScaleColor struct {
	Scale []*Struct_DamScaleColor_Scale `json:"scale"` // เกณฑ์สี
	Low   *Struct_DamScaleColor_LH      `json:"low"`   // ข้อความน้ำน้อย
	High  *Struct_DamScaleColor_LH      `json:"high"`  // ข้อความน้ำมาก
}

//	สร้างตัว Struct_DamScaleColor พร้อม setting ค่าจาก db
func New_Struct_DamScaleColor() (*Struct_DamScaleColor, error) {
	s := new(Struct_DamScaleColor)
	err := setting.GetSystemSettingPtr("Frontend.public.dam_scale_color", &s)
	if err != nil {
		return nil, err
	}
	return s, nil

}

//  เปรียบเทียบค่า v กับ struct Scale
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_DamScaleColor) CompareScale(v interface{}) *Struct_DamScaleColor_Scale {
	if s.Scale == nil {
		return nil
	}
	for _, sc := range s.Scale {
		if sc.Compare(v) {
			return sc
		}
	}
	return nil
}

type Struct_DamScaleColor_Scale struct {
	Operator  string `json:"operator"`  // example:`>` การกระทำ
	Term      string `json:"term"`      // example:`100` ค่า
	Color     string `json:"color"`     // example:`#C70000` สี
	Colorname string `json:"colorname"` // example:`min` ชื่อสี
	Text      string `json:"text"`      // example:`>100` ข้อความที่แสดง
	Level     string `json:"level"`     // example:`5` ระดับของสถานะ
}

//  เปรียบเทียบค่า v กับ struct
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_DamScaleColor_Scale) Compare(v interface{}) bool {
	b, err := Compare(v, s.Operator, s.Term)
	if err != nil {
		return false
	}
	return b
}

type Struct_DamScaleColor_LH struct {
	Color string `json:"color"` // example:`#FF0000` สี
	Trans string `json:"trans"` // example:`dam_high_text` ชื่อ trans ใน laravel
	Text  string `json:"text"`  // example:`<font color='#FF0000'>น้ำมาก(%รนก.)</font>` ข้อความที่แสดง

}

//
// Frontend.public.waterlevel_setting
//
type Struct_WaterlevelSetting struct {
	Scale      []*Struct_WaterlevelSetting_S `json:"scale"`      // เกณฑ์สี
	Rule       []*Struct_WaterlevelSetting_R `json:"rule"`       // เกณฑ์
	Level      *Struct_WaterlevelSetting_L   `json:"level"`      // การแสดงผลตามระดับ
	No_storage string                        `json:"no_storage"` // example:`` ข้อความที่จะมาแสดงตอนไม่มีค่าความจุลำน้ำ
	NotToday   *Struct_WaterlevelSetting_N   `json:"not_today"`  // ข้อมูลที่ไม่ใช่วันปัจจุบัน
}

//	สร้างตัว Struct_WaterlevelSetting พร้อม setting ค่าจาก db
func New_Struct_WaterlevelSetting() (*Struct_WaterlevelSetting, error) {
	s := new(Struct_WaterlevelSetting)
	err := setting.GetSystemSettingPtr("Frontend.public.waterlevel_setting", &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

//  เปรียบเทียบค่า v กับ struct Scale
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_WaterlevelSetting) CompareScale(v interface{}) *Struct_WaterlevelSetting_S {
	if s.Scale == nil {
		return nil
	}
	for _, sc := range s.Scale {
		if sc.Compare(v) {
			return sc
		}
	}
	return nil
}

type Struct_WaterlevelSetting_S struct {
	Operator  string `json:"operator"`  // example:`>` การกระทำ
	Term      string `json:"term"`      // example:`100` ค่า
	Color     string `json:"color"`     // example:`#FF0000` สี
	Colorname string `json:"colorname"` // example:`red` ชื่อสี
	Situation string `json:"situation"` // example:`น้ำล้นตลิ่ง` สถานการณ์
	Trans     string `json:"trans"`     // example:`waterlevel_level_5` ชื่อ trans ใน laravel
	Text      string `json:"text"`      // example:`>100` ข้อความที่แสดง
}

//  เปรียบเทียบค่า v กับ struct
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_WaterlevelSetting_S) Compare(v interface{}) bool {
	b, err := Compare(v, s.Operator, s.Term)
	if err != nil {
		return false
	}
	return b
}

type Struct_WaterlevelSetting_R struct {
	Operator string `json:"operator"` // example:`>` การกระทำ
	Term     string `json:"term"`     // example:`100` ค่า
	Level    int    `json:"level"`    // exmample:`2` ระดับ
}

type Struct_WaterlevelSetting_L struct {
	L1 *Struct_WaterlevelSetting_Level `json:"1"` // การแสดงผลตามระดับ
	L2 *Struct_WaterlevelSetting_Level `json:"2"` // การแสดงผลตามระดับ
}
type Struct_WaterlevelSetting_Level struct {
	Color     string `json:"color"`     // example:`#FF0000` สี
	Colorname string `json:"colorname"` // example:`red` ชื่อสี
	Trans     string `json:"trans"`     // example:`waterlevel_level_5` ชื่อ trans ใน laravel
}

type Struct_WaterlevelSetting_N struct {
	Color     string `json:"color"`     // example:`gray` สี
	Colorname string `json:"colorname"` // example:`gray` ชื่อสี
	Text      string `json:"text"`      // example:`-` ข้อความที่จะมาแสดงแทน
}

//
// Frontend.public.waterquality_setting
//
type Struct_WaterqualitySetting struct {
	Scale    *Struct_WaterqualitySetting_S `json:"scale"`     // เกณฑ์
	Criteria *Struct_WaterqualitySetting_C `json:"criteria"`  // เกณฑ์พิเศษ
	Default  *Struct_WaterqualitySetting_D `jons:"default"`   // ปกติ
	NotToDay *Struct_WaterqualitySetting_N `json:"not_today"` // ข้อมูลที่ไม่ใช่วันปัจจุบัน
}

type Struct_WaterqualitySetting_S struct {
	Salinity     []*Struct_WaterqualitySetting_Scale `json:"salinity,omitempty"`     // ความเค็ม
	Do           []*Struct_WaterqualitySetting_Scale `json:"do,omitempty"`           // ออกซิเจน
	PH           []*Struct_WaterqualitySetting_Scale `json:"ph,omitempty"`           // กรด-ด่าง
	Turbid       []*Struct_WaterqualitySetting_Scale `json:"turbid,omitempty"`       // ความขุ่นในน้ำ
	Tds          []*Struct_WaterqualitySetting_Scale `json:"tds,omitempty"`          // tds
	Conductivity []*Struct_WaterqualitySetting_Scale `json:"conductivity,omitempty"` // ความนำไฟฟ้าในน้ำ
}

type Struct_WaterqualitySetting_Scale struct {
	Operator string `json:"operator"` // example:`>` การกระทำ
	Term     string `json:"term"`     // example:`2` ค่า
	Color    string `json:"color"`    // example:`#EE141F` สี
	Text     string `json:"text"`     // example:`> 2` ข้อความที่แสดง
	Trans    string `json:"trans"`    // example:`salinity_2` ชื่อ trans ใน laravel
	InGraph  bool   `json:"inGraph"`  // example:`true` แสดงเส้นเกณฑ์ในกราฟ
}

type Struct_WaterqualitySetting_C struct {
	Struct_WaterqualitySetting_S
	Id   int64  `json:"id"`   // example:`144` รหัสของสถานีคุณภาพน้ำ
	Name string `json:"name"` // example:`สำเหร่` ชื่อสถานีคุณภาพน้ำ
}

type Struct_WaterqualitySetting_D struct {
	Color     string `json:"color"`     // example:`#66C803` สี
	Colorname string `json:"colorname"` // example:`green` ชื่อสี
}
type Struct_WaterqualitySetting_N struct {
	Color     string `json:"color"`     // example:`gray` สี
	Colorname string `json:"colorname"` // example:`gray` ชื่อสี
	Text      string `json:"text"`      // example:`-` ข้อความที่จะมาแสดงแทน
}

//Old Code
//type Arr_Struct_StormSetting struct {
//	Setting []*Struct_StormSetting
//}

type Arr_Struct_StormSetting []*Struct_StormSetting

//
// Frontend.public.storm_setting
//
type Struct_StormSetting struct {
	Operator  string `json:"operator"`   // example:`>` การกระทำ
	Term      string `json:"term"`       // example:`135` ค่า
	Color     string `json:"color"`      // example:`#CC00CC` สี
	Knots     string `json:"knots_text"` // example:`>135` knots
	Mph       string `json:"mph_text"`   // example:`>155` mph
	Kmh       string `json:"kmh_text"`   // example:`>250` km/h
	Category  string `json:"category"`   // example:`Cat 5` ประเภท
	Strength  string `json:"storm_type"` // example:`Typhoon Cat 5`
	ScaleText string `json:"scale_text"` // example:`รุนแรงมาก` ข้อความ
	Level     int    `json:"level"`      // example: 7
}

//	สร้างตัว Struct_DamScaleColor พร้อม setting ค่าจาก db
func Set_Struct_StormSetting() (Arr_Struct_StormSetting, error) {
	s := make(Arr_Struct_StormSetting, 0)
	err := setting.GetSystemSettingPtr("Frontend.public.storm_setting", &s)
//	sa := &Arr_Struct_StormSetting{}
//	sa.Setting = s
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (s Arr_Struct_StormSetting) CompareSetting(v interface{}) *Struct_StormSetting {
	for _, sc := range s {
		if sc.Compare(v) {
			return sc
		}
	}
	return nil
}

//  เปรียบเทียบค่า v กับ struct
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_StormSetting) Compare(v interface{}) bool {
	b, err := Compare(v, s.Operator, s.Term)
	if err != nil {
		return false
	}
	return b
}

//
// Frontend.public.warning_setting
//
type Struct_WarningSetting struct {
	Rain    string `json:"rain"`    // example:`#3c8dbc` สีของพื้นที่ประสบอุทกภัยและดินโคลนถล่ม
	Drought string `json:"drought"` // example:`#db802b` สีของภัยแล้ง
	Warning string `json:"warning"` // example:`#3c8dbc` สีขอพื้นที่ประกาศภัย
	Flood   string `json:"flood"`   // example:`#3c8dbc` สีของอุทกภัย
}

//
// Frontend.public.pre_rain_setting
//
type Struct_PreRainSetting struct {
	Rule  *Struct_PreRainSetting_R `json:"rule"`       // เกณฑ์แยกตามรหัสภาค
	Level *Struct_PreRainSetting_L `json:"level-text"` // แยกตามระดับ
}

type Struct_PreRainSetting_R struct {
	A []*Struct_PreRainSetting_Rain `json:"1"` // เกณฑ์แยกตามรหัสภาค
}
type Struct_PreRainSetting_Rain struct {
	Operator string `json:"operator"` // example:`>=` การกระทำ
	Term     string `json:"term"`     // example:`71` ค่า
	Level    int    `json:"level"`    // exmample:`1` ระดับ
}

type Struct_PreRainSetting_L struct {
	A []*Struct_PreRainSetting_Level `json:"1"` // ระดับ
}
type Struct_PreRainSetting_Level struct {
	Colorname string `json:"colorname"` // example:`red` ชื่อสี
	Color     string `json:"color"`     // example:`#ff0000` สี
	Trans     string `json:"trans"`     // example:`predict_rain_level_2` ชื่อ trans ใน laravel
	Text      string `json:"text"`      // example:`ฝนตกหนักมาก` ข้อความที่แสดง
}

//
// Frontend.public.wave_setting
//
type Arr_Struct_WaveSetting []*Struct_WaveSetting

func New_Struct_WaveSetting() (Arr_Struct_WaveSetting, error) {
	arr := make(Arr_Struct_WaveSetting, 0)
	err := setting.GetSystemSettingPtr("Frontend.public.wave_setting", &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

//  เปรียบเทียบค่า v กับ wave setting
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		Struct_WarningSetting
func (s Arr_Struct_WaveSetting) CompareSetting(v interface{}) *Struct_WaveSetting {
	for _, ss := range s {
		if ss.Compare(v) {
			return ss
		}
	}
	return nil
}

type Struct_WaveSetting struct {
	Operator  string `json"operator"`   // example:`>` การกระทำ
	Term      string `json:"term"`      // example:`2` ค่า
	Level     string `json:"level"`     // example:`3` ระดับ
	Status    string `json:"status"`    // example:`คลื่น สูงมากกว่า 2 เมตร` สถานะ
	Colorname string `json:"colorname"` // example:`red` ชื่อสี
	Color     string `json:"color"`     // example:`#ff0000` สี
	Trans     string `json:"trans"`     // example:`predict_wave_text` ชื่อ trans ใน laravel
}

//  เปรียบเทียบค่า v กับ struct
//	Parameters:
//		v
//			ค่าที่ต้องการเปรียบเทียบ
//	Return:
//		true ถ้าถูกเงื่อนไข
func (s *Struct_WaveSetting) Compare(v interface{}) bool {
	b, err := Compare(v, s.Operator, s.Term)
	if err != nil {
		return false
	}
	return b
}

//
// Frontend.public.dam_data_type
//
type Struct_DamDataType struct {
	Id    string           `json:"id"`    // example:`1` running number
	Value string           `json:"value"` // example:`dam_storage` ชื่อที่ต้องใช้ในโปรแกรม
	Text  *json.RawMessage `json:"text"`  // example:`{"th":"ปริมาตรกักเก็บ","en":"dam storage"}` ชื่อที่ต้องใช้แสดงบนเว็ป
}

//
// Frontend.public.rainfall_data_type
//
type Struct_RainfallDataType struct {
	Id    string           `json:"id"`    // example:`1` running number
	Value string           `json:"value"` // example:`dam_storage` ชื่อที่ต้องใช้ในโปรแกรม
	Text  *json.RawMessage `json:"text"`  // example:`{"th":"ปริมาณฝน 24 ชั่วโมงย้อนหลัง","en":"rain fall 24 hr latest"}` ชื่อที่ต้องใช้แสดงบนเว็ป
}

//
// Frontend.public.waterquality_data_type
//
type Struct_WaterqulityDataType struct {
	Id    string           `json:"id"`     // example:`1` running number
	Value string           `json:"value"`  // example:`salinity` ชื่อที่ต้องใช้ในโปรแกรม
	Text  string           `json:"string"` // example:`ความเค็ม(g/L)` ชื่อที่ต้องใช้แสดงบนเว็ป
	Name  *json.RawMessage `json:"name"`   // example:`{"th":"ความเค็ม(g/L)","en":"salinity(g/L)"}` ชื่อที่ต้องใช้แสดงบนเว็ปแบบหลายภาษา
}

type StructSeaWaterlevel struct {
	Name      string `json:"name"`      //example:`ตรวจวัดจริง` ชื่อข้อมูล
	Color     string `json:"color"`     //example:`#FF0000` รหัสสี
	ColorName string `json:"colorname"` //example:`red` ชื่อสี
}
