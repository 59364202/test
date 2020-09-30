package forecast

import (
	"encoding/json"
	waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	tw30_setting "haii.or.th/api/thaiwater30/util/setting"
)

type FloodforecastOutputWithScale struct {
	FloodforecastLoad []*FloodForecastOutPutCpy             `json:"floodforecast"`
	WaterlevelObserve []*waterlevel.ObserveWaterlevelOutput `json:"observe"`
	Scale             json.RawMessage                       `json:"scale"`
}

type FloodforecastOutputWithScaleSwagger struct {
	FloodforecastLoad []*FloodForecastOutPutCpy             `json:"floodforecast"`
	WaterlevelObserve []*waterlevel.ObserveWaterlevelOutput `json:"observe"`
	Scale             tw30_setting.Struct_WaterlevelSetting `json:"scale"`
}

type FloodForecastOutPutCpy struct {
	ID                    int64                    `json:"id"`                     // example:`134` รหัสข้อมูลคาดการณ์น้ำท่วม
	FloodForecastDatetime string                   `json:"floodforecast_datetime"` // example:`2006-01-02 15:04`วันเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FloodForecastValue    string                   `json:"floodforecast_value"`    // example:`15`ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
	Station               *FloodForecastStation    `json:"station"`                // สถานีคาดการณ์น้ำท่วม
	FloodForecastData     []*FloodForecastCpyValue `json:"data"`                   // ข้อมูลคาดการณ์น้ำท่วม
	Agency                *Agency                  `json:"agency"`                 // หน่วยงาน
}

type FloodForecastStation struct {
	ID                   int64           `json:"id"`                    // example:`134`รหัสสถานีคาดการณ์น้ำท่วม
	FloodforecastName    json.RawMessage `json:"floodforecast_name"`    // example:`{"th":"สถานีโทรมาตร"}`ชื่อสถานีโทรมาตร
	FloodForecastOldCode string          `json:"floodforecast_oldcode"` // example:`DD12`รหัสโทรมาตรเดิม
	Lat                  string          `json:"floodforecast_lat"`     // example:`9.141566`ละติจูดของสถานีโทรมาตร
	Long                 string          `json:"floodforecast_long"`    // example:`102.345123`ลองติจูดของสถานีโทรมาตร
	Alarm                interface{}     `json:"floodforecast_alarm"`   // example:`80`ระดับเตือนภัย
	Warning              interface{}     `json:"floodforecast_warning"` // example:`80` ระดับเตือนภัย
}

type Agency struct {
	ID        int64           `json:"id"`               // example:`9`รหัสหน่วยงาน
	Name      json.RawMessage `json:"agency_name"`      // example:`"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร"` ชื่อของหน่วยงาน (ภาษาไทย)
	ShortName json.RawMessage `json:"agency_shortname"` // example:`"th:"สสนก."`ชื่อหน่วยงาน
}

type FloodForecastCpyValue struct {
	FloodForecastDatetime string `json:"datetime"` // example:`2006-01-02 15:04`วันเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FloodForecastValue    string `json:"value"`    // example:`70`ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

type SwanStationOutput struct {
	SwanName json.RawMessage `json:"swan_name"` // example:`{"th":"สถานีโทรมาตร"}` ชื่อสถานีโทรมาตร
	Lat      string          `json:"swan_lat"`  // example:`9.231561` ละติจูดของสถานีโทรมาตร
	Long     string          `json:"swan_long"` // example:`103.242423` ลองติจูดของสถานีโทรมาตร
}

type SwanStationOutput2 struct {
	ID       int64           `json:"id"`        // example:`134` รหัสสถานีสถานี
	SwanName json.RawMessage `json:"swan_name"` // example:`{"th":"สถานีโทรมาตร"}` ชื่อสถานีโทรมาตร
	Lat      string          `json:"swan_lat"`  // example:`9.231561` ละติจูดของสถานีโทรมาตร
	Long     string          `json:"swan_long"` // example:`103.242423` ลองติจูดของสถานีโทรมาตร
}

type SwanForecastOutput struct {
	Station          *SwanStationOutput2  `json:"station"`            // สถานี
	SwanForecastData []*SwanForecastValue `json:"wave_forecast_data"` // ข้อมูลความการณ์
}

type SwanForecastValue struct {
	ID            int64       `json:"id"`                  // example:`2131`รหัสข้อมูลคาดการณ์
	Datetime      string      `json:"swan_datetime"`       // example:`2006-01-02`วันเวลาเก็บข้อมูลคาดการณ์ความสูงคลื่น
	Depth         interface{} `json:"swan_depth"`          // example:`65.4`ข้อมูลคาดการณ์ความลึกของคลื่น หน่วย m
	Highsig       interface{} `json:"swan_highsig"`        // example:`231.4`ข้อมูลคาดการณ์ความสูงของคลื่น หน่วย m
	Direction     interface{} `json:"swan_direction"`      // example:`51.4`ข้อมูลคาดการณ์ทิศทางของคลื่น หน่วย degree
	PeriodTop     interface{} `json:"swan_period_top"`     // example:`21.3`ข้อมูลคาดการณ์คาบคลื่นสูงสุด หน่วย sec
	PeriodAverage interface{} `json:"swan_period_average"` // example:`23.4`ข้อมูลคาดการณ์คาบคลื่นเฉลี่ย หน่วย sec
	WindX         interface{} `json:"swan_windx"`          // example:`3.1`ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศตะวันออกและทิศตะวันตก หน่วย m/s
	WindY         interface{} `json:"swan_windy"`          // example:`12.1`ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศเเหนือและใต้ หน่วย m/s
}

type Struct_Swan_Forecast_InputParam struct {
	StationID   string `json:"station_id"`   // required:false example:`1` รหัสสถานีคาดการณ์คลื่น ไม่ใส่ = ทุกสถานี เลือกได้หลายสถานี เช่น 1,2,3
}

type FloodForecastMonitoring struct {
	SubbasinName      json.RawMessage `json:"subbasin_name"`          // example:`{"th":"ลุ่มน้ำสาขา"}`ชื่อลุ่มน้ำสาขา
	FloodforecastName json.RawMessage `json:"floodforecast_name"`     // example:`{"th":"สถานีโทรมาตร"}`ชื่อสถานีโทรมาตร
	Datetime          string          `json:"floodforecast_datetime"` // example:`2006-01-02 15:04`วันเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	Value             interface{}     `json:"floodforecast_value"`    // example:`1`ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}
