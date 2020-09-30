package cctv

import (
	"encoding/json"

	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

type cctvOutput struct {
	Id              int64           `json:"id"`                          // example:`1` รหัส CCTV
	DamID           string          `json:"dam_id"`                      // example:`26` รหัสเขื่อน id จากตาราง m_dam
	TeleStationName json.RawMessage `json:"tele_station_name,omitempty"` // example:`stationid` รหัสสถานีโทรมาตร
	BasinName       json.RawMessage `json:"basin_name,omitempty"`        // example:`10`รหัสลุ่มน้ำ
	Lat             string          `json:"lat"`                         // example:`17.244115`พิกัดละติจูด
	Long            string          `json:"long"`                        // example:`98.972687`พิกัดลองติจูด
	SubBasin_id     int64           `json:"sub_basin_id,omitempty"`      // example:`21` รหัสลุ่มน้ำย่อย
	Title           string          `json:"title"`                       // example:`เขื่อนภูมิพล`ชื่อสถานที่
	Description     string          `json:"description"`                 // example:`description`คำอธิบายสถานที่
	MediaType       string          `json:"media_type"`                  // example:`img`ชนิดไฟล์วิดีโอ
	URL             string          `json:"cctv_url"`                    // example:`http://cctv1.bhumiboldam.egat.com/` ที่อยู่วิดีโอ
	CctvFlash       string          `json:"cctv_flash"`                  // example:`<EMBED src="http://203.150.226.24/flowplayer/flowplayer.swf" quality="high" menu="true" pluginspage="http://www.macromedia.com/shockwave/download/index.cgi?P1_Prod_Version=ShockwaveFlash" type="application/x-shockwave-flash" width="100%" height="100%" flashvars="config=%7B%22clip%22%3A%7B%22provider%22%3A%22rtmp%22%2C%22live%22%3Atrue%2C%22url%22%3A%22cpy2.stream%22%2C%22autoPlay%22%3Atrue%7D%2C%22playlist%22%3A%5B%7B%22provider%22%3A%22rtmp%22%2C%22live%22%3Atrue%2C%22url%22%3A%22cpy2.stream%22%2C%22autoPlay%22%3Atrue%7D%5D%2C%22plugins%22%3A%7B%22controls%22%3A%7B%22time%22%3Afalse%2C%22volume%22%3Afalse%2C%22mute%22%3Afalse%2C%22autoHide%22%3Atrue%2C%22scrubber%22%3Afalse%2C%22url%22%3A%22http%3A//203.150.226.24/flowplayer/flowplayer.controls.swf%22%7D%2C%22rtmp%22%3A%7B%22url%22%3A%22http%3A//203.150.226.24/flowplayer/flowplayer.rtmp.swf%22%2C%22netConnectionUrl%22%3A%22rtmp%3A//203.150.226.24/rtplive%22%7D%2C%22sharing%22%3A%7B%22url%22%3A%22http%3A//203.150.226.24/flowplayer/flowplayer.sharing.swf%22%2C%22buttons%22%3A%7B%22overColor%22%3A%22%23ff0000%22%7D%2C%22share%22%3A%7B%22shareUrl%22%3A%22http%3A//203.150.226.24/%22%7D%7D%2C%22dock%22%3A%7B%22horizontal%22%3Afalse%2C%22right%22%3A10%2C%22width%22%3A%2210pct%22%2C%22top%22%3A5%7D%7D%7D" bgcolor="#000000" quality="true"> </EMBED>"> </EMBED>` ที่อยู่วิดีโอ
	CctvQuickTime   string          `json:"cctv_quicktime"`              // example:`<embed src="http://203.150.226.24:1935/live/makhamtao.stream/playlist.m3u8" width="100%" height="100%" autoplay="true" scale="tofit" controller="true" pluginspage="http://www.apple.com/quicktime/">` ที่อยู่วิดีโอ
	IsActive        bool            `json:"is_active"`                   // example: true

	Geocode *model_lt_geocode.Struct_Geocode `json:"geocode"` // ข้อมูลขอบเขตการปกครองของประเทศไทย
	Basin   *model_basin.Struct_Basin        `json:"basin"`   // ลุ่มน้ำ
}
