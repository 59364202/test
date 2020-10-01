package tele_waterlevel

import (
	"database/sql"
	"encoding/json"
	"fmt"

	uDatetime "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/float"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/pqx"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"

	uSetting "haii.or.th/api/thaiwater30/util/setting"
)

type Struct_Waterlevel struct {
	Id                  int64       `json:"id,omitempty"`            // example:`68033565` รหัสข้อมูล
	Waterlevel_datetime string      `json:"waterlevel_datetime"`     // example:`2006-01-02 15:04` วันเวลาที่ตรวจสอบค่าระดับน้ำ
	Waterlevel_m        interface{} `json:"waterlevel_m"`            // example:`0.5`ระดับน้ำ เมตร
	Waterlevel_msl      interface{} `json:"waterlevel_msl"`          // example:`167.953642` ระดับน้ำ รทก
	Pre_Waterlevel_msl  interface{} `json:"waterlevel_msl_previous"` // example:`167.246358` ระดับน้ำ รทก ครั้งก่อนหน้า
	FlowRate            interface{} `json:"flow_rate"`               // example:`28.62` อัตราการไหล
	Discharge           interface{} `json:"discharge"`               // example:`68.32` ปริมาณการระบายน้ำ
	Storage_percent     interface{} `json:"storage_percent"`         // example:`1325.5626182618266` % เทียบความจุลำน้ำ
	SortOrder           interface{} `json:"sort_order"`              // example:`null, 1` เรียงลำดับ
	Table               string      `json:"station_type"`            // example:`tele_waterlevel` ชนิดของสถานี

	Agency  *model_agency.Struct_Agency            `json:"agency"`  // หน่วยงาน
	Basin   *model_basin.Struct_Basin              `json:"basin"`   // ลุ่มน้ำ
	Station *model_tele_station.Struct_TeleStation `json:"station"` // สถานี
	Geocode *model_lt_geocode.Struct_Geocode       `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย

	//Datetime     string          `json:"datetime"`
	//Oldcode      string          `json:"oldcode"`
	//Value        interface{}     `json:"value"`
	//Name         json.RawMessage `json:"name,omitempty"`
	//ProvinceName json.RawMessage `json:"province_name,omitempty"`
	//DataID       int64           `json:"data_id"`
	//StationID    int64           `json:"station_id"`
}

type WaterlevelLastestStruct struct {
	Id             int64       `json:"id"`                        // example:`114497665` รหัสข้อมูลระดับน้ำ
	Station_Id     int64       `json:"tele_station_id,omitempty"` // example:`17` รหัสสถานี
	Datetime       string      `json:"waterlevel_datetime"`       // example:`2006-01-02 15:04` วันเวลาที่ตรวจสอบค่าระดับน้ำ
	Waterlevel_m   interface{} `json:"waterlevel_m"`              // example:`null` ระดับน้ำ เมตร
	Waterlevel_msl interface{} `json:"waterlevel_msl"`            // example:`129.05` ระดับน้ำ รทก
	Flow_rate      interface{} `json:"flow_rate"`                 // example:`98.17` อัตราการไหล
	Discharge      interface{} `json:"discharge"`                 // example:`null` ปริมาณการระบายน้ำ
}

type Waterlevel_InputParam struct {
	Id            string `json:"id"`
	Station_id    string `json:"station_id"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	Subbasin_id   string `json:"subbasin_id"`
	Agency_id     string `json:"agency_id"`
	Basin_id      string `json:"basin_id"`      // required:false example:`1` รหัสลุ่มน้ำ ไม่ใส่ = ทุกลุ่มน้ำ เลือกได้หลายลุ่มน้ำ เช่น 1,2,4
	Region_Code   string `json:"region_code"`   // required:false example:`1` รหัสภาค ไม่ใส่ = ทุกภาค เลือกได้ทีละภาค
	Province_Code string `json:"province_code"` // required:false example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด เลือกได้หลายจังหวัด เช่น 10,51,62
	Order         string `json:"order"`         // required:false example:`asc` การเรียงข้อมูล storage_percent asc= เรียงจากค่าน้อยไปหามาก desc = เรียงจากค่ามากไปหาน้อย
	Data_Limit    int    `json:"data_limit"`    // จำนวนข้อมูล

	IsHourly bool
	IsMain   bool `json:"-"`
}

type Struct_Waterlevel_ErrorData struct {
	ID              int64           `json:"id"`                    // รหัสข้อมูลระดับน้ำ
	StationID       int64           `json:"station_id"`            // รหัสสถานี
	StationOldCode  string          `json:"station_oldcode"`       // รหัสสถานีเดิม
	Datetime        string          `json:"datetime"`              // วันเวลา
	StationName     json.RawMessage `json:"station_name"`          // ชื่อสถานี
	ProvinceName    json.RawMessage `json:"station_province_name"` // ชื่อจังหวัด
	AgencyName      json.RawMessage `json:"agency_name"`           // ชื่อหน่วยงาน
	AgencyShortName json.RawMessage `json:"agency_shortname"`      // ชื่อย่อหน่วยงาน
	WaterlevelM     interface{}     `json:"waterlevel_m"`          // ระดับน้ำ เมตร
	WaterlevelMsl   interface{}     `json:"waterlevel_msl"`        // ระดับน้ำ รทก
	//WaterlevelIn    interface{}     `json:"waterlevel_in"`
	//WaterlevelOut   interface{}     `json:"waterlevel_out"`
	//WaterlevelOut2  interface{}     `json:"waterlevel_out2"`
	FlowRate  interface{} `json:"flow_rate"` // อัตราการไหล
	Discharge interface{} `json:"discharge"` // ปริมาณการระบายน้ำ
}

type GetWaterlevelLastest_OutputParam struct {
	Header []string                   `json:"header"` // example:`["id","waterlevel_datetime","waterlevel_m","waterlevel_msl","waterlevel_in","waterlevel_out","waterlevel_out2","flow_rate","discharge"]` หัวตาราง
	Data   []*WaterlevelLastestStruct `json:"data"`   // ข้อมูล
}

type GetWaterlevelGraphByStationAndDateAnalystDataOutput struct {
	Datetime string      `json:"datetime"` // example:`2006-01-02 15:04` วันเวลา
	Value    interface{} `json:"value"`    // example:`137.91` ระดับน้ำ รทก
}

type WaterlevelGDataOutput struct {
	GraphData   []*GetWaterlevelInOutGrapthAnalystOutput `json:"graph_data"`   // ข้อมูลกราฟ
	GroundLevel interface{}                              `json:"ground_level"` // ระดับท้องน้ำ ม.รทก
	MinBank     interface{}                              `json:"min_bank"`     // ตลิ่งต่ำสุด
}

type GetWaterlevelGraphByStationAndDateAnalystOutput struct {
	GraphData   []*GetWaterlevelGraphByStationAndDateAnalystDataOutput `json:"graph_data"`   // ข้อมูลกราฟ
	MinBank     interface{}                                            `json:"min_bank"`     // example:`0.881` ตลิ่งต่ำสุด
	GroundLevel interface{}                                            `json:"ground_level"` // example:`-5.368` ระดับท้องน้ำ ม.รทก
}

type GetWaterlevelYearlyGraphInput struct {
	StationID   string `json:"station_id"`   // รหัสสถานี
	StationType string `json:"station_type"` // ชนิดสถานี
	Year        []int  `json:"year"`         // ปี
}

type GetWaterlevelYearlyGraphAnalystOutput struct {
	GraphData   []*GetWaterlevelYearlyGraphAnalystOutputYear `json:"graph_data"`   // ข้อมูลกราฟ
	MinBank     interface{}                                  `json:"min_bank"`     // example:`1.13`ตลิ่งต่ำสุด
	GroundLevel interface{}                                  `json:"ground_level"` // example:`0.12`ระดับท้องน้ำ ม.รทก
}

type GetWaterlevelYearlyGraphAnalystOutputYear struct {
	Year      int                                                    `json:"year"`       // example:`2016`ปี
	GraphData []*GetWaterlevelGraphByStationAndDateAnalystDataOutput `json:"graph_data"` // ข้อมูลกราฟ
}

type GetWaterlevelInOutGrapthAnalystInput struct {
	StationID int64  `json:"station_id"` // รหัสสถานี เช่น 10,23,413
	StartDate string `json:"start_date"` // เวลาเริ่มต้น 2006-01-02
	EndDate   string `json:"end_date"`   // เวลาสิ้นสุด 2006-01-02
}

type GetWaterlevelInOutGrapthAnalystOutput struct {
	Name          string                                                 `json:"name"`           // example:`เหนือ ปตร.`ชื่อ
	Data          []*GetWaterlevelGraphByStationAndDateAnalystDataOutput `json:"data"`           // ข้อมูลกราฟ
	GroundLevel   interface{}                                            `json:"ground_level"`   // example:`0.12`ระดับท้องน้ำ ม.รทก
	CriticalLevel interface{}                                            `json:"critical_level"` // example:`1.59`ค่าระดับแถบเตือนภัยสูงสุด  ม.รทก
}

// for water gate ปตร./ฝาย
type GetWaterlevelInOutLatestAnalystOutput struct {
	DatetimeIn      string                                 `json:"watergate_datetime_in"`  // example:`2006-01-02 15:04` วันที่ตรวจสอบค่าระดับน้ำด้านในประตู
	WaterlevelIn    interface{}                            `json:"watergate_in"`           // example:`` ระดับน้ำด้านในประตูระบายน้ำ
	DatetimeOut     string                                 `json:"watergate_datetime_out"` // example:`2006-01-02 15:04` วันที่ตรวจสอบค่าระดับน้ำด้านนอกประตู
	WaterlevelOut   interface{}                            `json:"watergate_out"`          // example:`12.3`ระดับน้ำด้านนอกประตูระบายน้ำ
	WaterlevelOut2  interface{}                            `json:"watergate_out2"`         // example:`31.3`ระดับน้ำด้านนอกประตูระบายน้ำ
	PumpOn          interface{}                            `json:"pump_on"`                // example:`23`จำนวนเครื่องสูบน้ำที่ใช้งาน (เครื่อง)
	Pump            interface{}                            `json:"pump"`                   // example:`40`จำนวนเครื่องสูบน้ำ (เครื่อง)
	FloodgateOpen   interface{}                            `json:"floodgate_open"`         // example:`10` จำนวนประตูระบายน้ำที่เปิด (บาน)
	Floodgate       interface{}                            `json:"floodgate"`              // example:`13` จำนวนประตูระบายน้ำ (บาน)
	FloodgateHeight interface{}                            `json:"floodgate_height"`       // example:`23` จำนวนความสูงของประตูระบายน้ำ (เมตร)
	Agency          *model_agency.Struct_Agency            `json:"agency"`                 // หน่วยงาน
	Basin           *model_basin.Struct_Basin              `json:"basin"`                  // ลุ่มน้ำ
	Station         *model_tele_station.Struct_TeleStation `json:"station"`                // สถานี
	Geocode         *model_lt_geocode.Struct_Geocode       `json:"geocode"`                // ข้อมูลขอบเขตการปกครองของประเทศไทย
}

type Basin struct {
	Id         interface{}     `json:"id"`                   // example:`1` รหัสลุ่มน้ำ
	Basin_code int64           `json:"basin_code,omitempty"` // example:`1` รหัสลุ่มน้ำ
	Basin_name json.RawMessage `json:"basin_name,omitempty"` // example:`{"th":"ลุ่มน้ำสาละวิน","en":"MAE NAM SALAWIN"}` ชื่อลุ่มน้ำ
}

type Struct_onLoad_Basin struct {
	Result string                      `json:"result"` // example:`OK`
	Data   []*model_basin.Struct_Basin `json:"data"`   // ลุ่มน้ำ
}
type Struct_onLoad_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

type WaterlevelTeleStation struct {
	ID      int64              `json:"id"`                           // example:`145`รหัสสถานี
	Name    json.RawMessage    `json:"tele_station_name,ommitempty"` // example:`{"th":"สภานีระดับน้ำ"}`ชื่อสถานี
	Lat     interface{}        `json:"tele_station_lat"`             // example:`9.041243`ละติจูด
	Long    interface{}        `json:"tele_station_long"`            // example:`103.441444`ลองติจูด
	OldCode interface{}        `json:"tele_station_oldcode"`         // example:``รหัสสถานีเดิม
	Geocode *WaterlevelGeocode `json:"geocode,ommitempty"`           // ข้อมูลขอบเขตการปกครองของประเทศไทย
}

type WaterlevelGeocode struct {
	ID           interface{}      `json:"id"`                       // example:`3`ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
	Geocode      interface{}      `json:"geocode"`                  // example:`100101`รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
	ProvinceCode interface{}      `json:"province_code"`            // example:`10`รหัสจังหวัด
	ProvinceName *json.RawMessage `json:"province_name,ommitempty"` // example:`{"th":"กรุงเทพมหานคร"}`ชื่อจังหวัด
	AmphoeCode   interface{}      `json:"amphoe_code"`              // example:`01`รหัสอำเภอ
	AmphoeName   *json.RawMessage `json:"amphoe_name,ommitempty"`   // example:`{"th":"พระนคร"}`ชื่ออำเภอ
	TumbonCode   interface{}      `json:"tumbon_code"`              // example:`06`รหัสตำบล
	TumbonName   *json.RawMessage `json:"tumbon_name,ommitempty"`   // example:`{"th":"เสาชิงช้า"}`ชื่อตำบล
	AreaCode     interface{}      `json:"area_code"`                // example:`2`รหัสภาค
	//	AreaName     *json.RawMessage `json:"area_name,ommitempty"`     // example:`{"th":"ภาคกลาง"}`ชื่อภาค
	//	AreaCode string           `json:"area_code,omitempty"` // example:`1` รหัสภาคของประเทศไทย
	AreaName *json.RawMessage `json:"area_name,omitempty"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อภาคของประเทศไทย
}

var WaterlevelColumnForCheckErrordata = []string{
	"id",
	"station_oldcode",
	"datetime",
	"station_name",
	"station_province_name",
	"agency_name",
	"agency_shortname",
	"waterlevel_m",
	"waterlevel_msl",
	"waterlevel_in",
	"waterlevel_out",
	"waterlevel_out2",
	"flow_rate",
	"discharge",
}

type Struct_WaterlevelBasinGraphAnalystAdvance struct {
	Name        *json.RawMessage `json:"name"`         // example:`{"th":"บ้านห้วยแก้ว"}` ชื่อสถานี
	Value       interface{}      `json:"value"`        // example:`154.36` ระดับน้ำ รทก
	GroundLevel interface{}      `json:"ground_level"` // example:`136.54` ระดับท้องน้ำ ม.รทก
	MinBank     interface{}      `json:"min_bank"`     // example:`130.54` ตลิ่งตำสุด
}

type Struct_WaterlevelBasin24HGraphAnalystAdvance struct {
	Datetime  string                                                    `json:"datetime"`   // example:`2006-01-02` วันเวลา
	GraphData []*Struct_WaterlevelBasin24HGraphAnalystAdvance_GraphData `json:"graph_data"` // ข้อมูลกราฟ
}

type Struct_WaterlevelBasin24HGraphAnalystAdvance_GraphData struct {
	Name        *json.RawMessage `json:"name"`         // example:`{"th":"บ้านห้วยแก้ว"}` ชื่อสถานี
	Value       interface{}      `json:"value"`        // example:`154.36` ระดับน้ำ รทก
	GroundLevel interface{}      `json:"ground_level"` // example:`136.54` ระดับท้องน้ำ ม.รทก
	LeftBank    interface{}      `json:"left_bank"`    // example:`135.45` ระดับตลิ่ง (ซ้าย)
	RightBank   interface{}      `json:"right_bank"`   // example:`130.54` ระดับตลิ่ง (ขวา)
	//	Lat         interface{}      `json:"lat"`
	//	Long        interface{}      `json:"long"`
	Distance interface{} `json:"distance"` // example:`200` ระยะทางการไหลของน้ำระหว่างสถานี
}

type ObserveWaterlevelOutput struct {
	ID                     int64                  `json:"id"`                       // example:`1245`รหัสข้อมูลระดับน้ำ
	TeleWaterlevelDatetime string                 `json:"tele_waterlevel_datetime"` // example:`2006-01-02 15:04`วันเวลาที่ตรวจสอบค่าระดับน้ำ
	TeleWaterlevelValue    interface{}            `json:"tele_waterlevel_value"`    // example:`90`ระดับน้ำ รทก
	Station                *WaterlevelTeleStation `json:"station"`                  // สถานี
	Agency                 *Agency                `json:"agency"`                   // หน่วยงาน
}

type Agency struct {
	ID        int64           `json:"id"`               // examplel:`9`รหัสสถานี
	Name      json.RawMessage `json:"agency_name"`      // example:`"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร"` ชื่อของหน่วยงาน (ภาษาไทย)
	ShortName json.RawMessage `json:"agency_shortname"` // example:`"th:"สสนก."`ชื่อหน่วยงาน
}

type Struct_WaterlevelMinMax struct {
	Current_time        string `json:"current_time"`        // example:`12:56 น.` เวลา
	Station_name_max    string `json:"station_name_max"`    // exmaple:`สถานีคลองแสนแสบ บางกะปิ จ.กรุงเทพมหานคร` ชื่อสถานี
	Percent_max         string `json:"percent_max"`         // example:`80.30%` ระดับน้ำ (%)
	Wl_msl_max          string `json:"wl_msl_max"`          // example:`-0.35` ระดับน้ำ
	Current_date_thai   string `json:"current_date_thai"`   // example:`18 ต.ค. 2558` วันที่
	Wl_max_before_level string `json:"wl_max_before_level"` // example:`fa fa-arrow-up` สัญลักษณ์
	Wl_max_color        string `json:"wl_max_color"`        // example:`#003CFA` สี
	Station_name_min    string `json:"station_name_min"`    // example:`สถานีคลองลาดพร้าว ท้ายปตร.คลอง2 จ.กรุงเทพมหานคร` ชื่อสถานี
	Percent_min         string `json:"percent_min"`         // example:`30.92%` ระดับน้ำ (%)
	Wl_msl_min          string `json:"wl_msl_min"`          // example:`-0.52` ระดับน้ำ
	Wl_min_before_level string `json:"wl_min_before_level"` // example:`fa fa-arrow-down` สัญลักษณ์
	Wl_min_color        string `json:"wl_min_color"`        // example:`#00B050` สี
}

type temp_WaterLevelMinMax struct {
	Time    string  // เวลา
	Name    string  // ชื่อสถานี
	Percent string  // ระดับน้ำ (%)
	Wl      string  // ระดับน้ำ
	Date    string  // วันที่
	Level   string  // สัญลักษณ์
	Color   string  // สี
	Value   float64 // ระดับน้ำ (%)
}

//	สร้าง struct temp_WaterLevelMinMax จาก sql.row
func newTemp_WaterLevelMinMax(row *sql.Rows) (*temp_WaterLevelMinMax, error) {
	var (
		_datetime_current  sql.NullString
		_value_current     sql.NullFloat64
		_value_previous    sql.NullFloat64
		_storage_percent   sql.NullFloat64
		_tele_station_name sql.NullString
		_province_name     sql.NullString
	)

	err := row.Scan(&_datetime_current, &_value_current, &_value_previous, &_storage_percent, &_tele_station_name, &_province_name)
	if err != nil {
		return nil, err
	}

	datetime := pqx.NullStringToTime(_datetime_current)
	st := new(temp_WaterLevelMinMax)

	st.Time = datetime.Format("15:04 น.")
	st.Name = _tele_station_name.String + " จ." + _province_name.String
	st.Percent = float.String(_storage_percent.Float64, 2) + " %"
	st.Wl = fmt.Sprintf("%.2f", _value_current.Float64)
	st.Date = datetime.Format("02 ") + uDatetime.MonthTHShort(datetime.Month()) + " " + datatype.MakeString(datetime.Year()+543)
	st.Value = _storage_percent.Float64

	// เทียบ ค่าปัจจุบัน กับ ค่าก่อนหน้า
	if _value_current.Float64 > _value_previous.Float64 {
		st.Level = "fa fa-arrow-up" // ^
	} else if _value_current.Float64 == _value_previous.Float64 {
		st.Level = "fa fa-circle-o" // o
	} else {
		st.Level = "fa fa-arrow-down" // v
	}

	return st, nil
}

//	Compare ค่า wl กับตัว setting
func (st *temp_WaterLevelMinMax) CompareScale(setting uSetting.Struct_WaterlevelSetting) error {
	return nil
}
