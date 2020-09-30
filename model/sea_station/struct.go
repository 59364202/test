package sea_station

import (
	"encoding/json"
)

type Agency struct {
	ID        int64           `json:"id"`               // example:`1` รหัสหน่วยงาน
	Name      json.RawMessage `json:"agency_name"`      // example:`{"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร"}` ชื่อหน่วยงาน
	ShortName json.RawMessage `json:"agency_shortname"` // example:`{"th":"สสนก."}` ชื่อย่อหน่วยงาน
	Station   []*Station      `json:"sea_station"`      // สถานี
}

type Station struct {
	ID      int64           `json:"id"`                  // example:`1` รหัสสถานี
	Name    json.RawMessage `json:"sea_station_name"`    // example:`{"th":"สถานี ก."}`ชื่อสถานี
	Lat     string          `json:"sea_station_lat"`     // example:`9.112311`ละติจูดของสถานี
	Long    string          `json:"sea_station_long"`    // example:`100.044242` ลองติจูดของสถานี
	Oldcode string          `json:"sea_station_oldcode"` // example:`AA03`รหัสสถานีเดิม
}
